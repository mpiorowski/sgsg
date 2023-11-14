import { getFormValue } from "$lib/helpers";
import { grpcSafe } from "$lib/safe";
import { upsendApi } from "$lib/server/api";
import { server } from "$lib/server/grpc";
import { createMetadata } from "$lib/server/metadata";
import { error, fail } from "@sveltejs/kit";

/** @type {import('./$types').PageServerLoad} */
export async function load({ locals }) {
    /**
     * We return the profile data immediately, and then fetch the resume and stream it to the client as it loads.
     */

    /** @type {import('$lib/safe').Safe<import('$lib/proto/proto/Profile').Profile__Output>} */
    const profile = await new Promise((r) => {
        server.GetProfile({}, createMetadata(locals.token), grpcSafe(r));
    });
    if (profile.error) {
        throw error(500, profile.msg);
    }

    /** @type {Promise<{
     *  name: string;
     *  base64: string;
     *  mimeType: string;
     * }>}
     */
    const resumePromise = new Promise((resolve, reject) => {
        let resume = {
            name: "",
            base64: "",
            mimeType: "",
        };
        if (profile.data.resumeId) {
            /** @type {import("$lib/safe").Safe<import("$lib/types").UpsendFile>} */
            upsendApi({
                url: `/files/server/${profile.data.resumeId}`,
                method: "GET",
            }).then((res) => {
                if (res.error) {
                    reject(res.msg);
                } else {
                    resume = {
                        name: res.data.name,
                        base64: Buffer.from(res.data.buffer).toString("base64"),
                        mimeType: res.data.mime_type,
                    };
                    resolve(resume);
                }
            });
        } else {
            resolve(resume);
        }
    });

    return {
        profile: profile.data,
        stream: { resume: resumePromise },
    };
}

/** @type {import('./$types').Actions} */
export const actions = {
    createProfile: async ({ locals, request }) => {
        const form = await request.formData();

        let resumeId = getFormValue(form, "resumeId");
        const resume = form.get("resume");
        if (!(resume instanceof File)) {
            return fail(400, { error: "Resume must be a PDF" });
        }
        if (resume.size > 0) {
            if (resume.size > 5 * 1024 * 1024) {
                return fail(400, { error: "Resume must be less than 5MB" });
            }
            if (!resume.name.endsWith(".pdf")) {
                return fail(400, { error: "Resume must be a PDF" });
            }

            // Upload new resume
            /** @type {import("$lib/safe").Safe<import("$lib/types").UpsendFile>} */
            const file = await upsendApi({
                url: "/files",
                method: "POST",
                file: resume,
            });
            if (file.error) {
                return fail(400, { error: file.msg });
            }

            // Delete old resume
            if (resumeId) {
                const resDel = await upsendApi({
                    url: `/files/${resumeId}`,
                    method: "DELETE",
                });
                if (resDel.error) {
                    return fail(400, { error: resDel.msg });
                }
            }
            resumeId = file.data.id;
        }

        let coverId = getFormValue(form, "coverId");
        let coverUrl = getFormValue(form, "coverUrl");
        const cover = form.get("cover");
        if (!(cover instanceof File)) {
            return fail(400, { error: "Cover must be an image" });
        }
        if (cover.size > 0) {
            if (cover.size > 5 * 1024 * 1024) {
                return fail(400, { error: "Cover must be less than 5MB" });
            }
            const extensions = [".png", ".jpg", ".jpeg", ".gif", ".svg"];
            if (!extensions.some((ext) => cover.name.endsWith(ext))) {
                return fail(400, { error: "Cover must be an image" });
            }

            // Upload new cover
            /** @type {import("$lib/safe").Safe<import("$lib/types").UpsendImage>} */
            const file = await upsendApi({
                url: "/images",
                method: "POST",
                file: cover,
            });
            if (file.error) {
                return fail(400, { error: file.msg });
            }

            // Delete old cover
            if (coverId) {
                const resDel = await upsendApi({
                    url: `/images/${coverId}`,
                    method: "DELETE",
                });
                if (resDel.error) {
                    return fail(400, { error: resDel.msg });
                }
            }
            coverId = file.data.id;
            coverUrl = file.data.url;
        }

        /** @type {import('$lib/proto/proto/Profile').Profile} */
        const data = {
            id: getFormValue(form, "id"),
            username: getFormValue(form, "username"),
            about: getFormValue(form, "about"),
            resumeId: resumeId,
            coverId: coverId,
            coverUrl: coverUrl,
        };

        /** @type {import("$lib/safe").Safe<import("$lib/proto/proto/Profile").Profile__Output>} */
        const res = await new Promise((r) => {
            server.CreateProfile(
                data,
                createMetadata(locals.token),
                grpcSafe(r),
            );
        });

        if (res.error) {
            if (res.fields) {
                return fail(400, { fields: res.fields });
            }
            return fail(400, { error: res.msg });
        }

        return {
            profile: res.data,
        };
    },
};