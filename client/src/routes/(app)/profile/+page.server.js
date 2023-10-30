import { getFormValue } from "$lib/helpers";
import { grpcSafe } from "$lib/safe";
import { upsendApi } from "$lib/server/api";
import { server } from "$lib/server/grpc";
import { createMetadata } from "$lib/server/metadata";
import { error, fail } from "@sveltejs/kit";

/** @type {import('./$types').PageServerLoad} */
export async function load({ locals }) {
    /** @type {import('$lib/safe').Safe<import('$lib/proto/proto/Profile').Profile__Output>} */
    const res = await new Promise((r) => {
        server.GetProfile({}, createMetadata(locals.token), grpcSafe(r));
    });
    if (res.error) {
        throw error(500, res.msg);
    }
    return {
        profile: res.data,
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
            /** @type {import("$lib/safe").Safe<import("$lib/types").UpsendFile>} */
            const file = await upsendApi({
                url: "/files",
                method: "POST",
                file: resume,
            });
            if (file.error) {
                return fail(400, { error: file.msg });
            }
            resumeId = file.data.id;
        }

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
            /** @type {import("$lib/safe").Safe<import("$lib/types").UpsendImage>} */
            const file = await upsendApi({
                url: "/images",
                method: "POST",
                file: cover,
            });
            if (file.error) {
                return fail(400, { error: file.msg });
            }
            coverUrl = file.data.url;
        }

        /** @type {import('$lib/proto/proto/Profile').Profile} */
        const data = {
            id: getFormValue(form, "id"),
            username: getFormValue(form, "username"),
            about: getFormValue(form, "about"),
            resumeId: resumeId,
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
