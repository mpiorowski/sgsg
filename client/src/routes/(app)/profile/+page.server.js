import { getFormValue } from "$lib/helpers";
import { safe } from "$lib/server/safe";
import { grpcSafe } from "$lib/server/grpc";
import { upsendApi } from "$lib/server/api";
import { server } from "$lib/server/grpc";
import { perf } from "$lib/server/logger";
import { createMetadata } from "$lib/server/metadata";
import { error, fail } from "@sveltejs/kit";

/** @type {import('./$types').PageServerLoad} */
export async function load({ locals }) {
    const end = perf("load_profile");
    /**
     * We return the profile data immediately, and then fetch the resume and stream it to the client as it loads.
     */

    /** @type {import('$lib/server/safe').Safe<import('$lib/proto/proto/Profile').Profile__Output>} */
    const profile = await new Promise((r) => {
        server.GetProfileByUserId(
            {},
            createMetadata(locals.user.id),
            grpcSafe(r),
        );
    });
    if (!profile.success) {
        throw error(500, profile.error);
    }

    /** @type {Promise<import("$lib/server/safe").Safe<import("$lib/types").UpsendFile | undefined>>} */
    let resumePromise = Promise.resolve({ data: undefined, success: true });
    if (profile.data.resume_id) {
        /** @type {Promise<import('$lib/server/safe').Safe<import('$lib/types').UpsendFile>>} */
        resumePromise = upsendApi({
            url: `/files/${profile.data.resume_id}`,
            method: "GET",
        });
    }

    end();
    return {
        profile: profile.data,
        stream: { resume: resumePromise },
    };
}

/** @type {import('./$types').Actions} */
export const actions = {
    createProfile: async ({ locals, request }) => {
        const end = perf("create_profile");
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

            /**
             * Delete old resume
             */
            if (resumeId) {
                const resDel = await upsendApi({
                    url: `/files/${resumeId}`,
                    method: "DELETE",
                });
                if (!resDel.success) {
                    return fail(400, { error: resDel.error });
                }
            }

            /**
             * Upload new resume
             * @type {import("$lib/server/safe").Safe<import("$lib/types").UpsendFile>}
             */
            const file = await upsendApi({
                url: "/files",
                method: "POST",
                file: resume,
            });
            if (!file.success) {
                return fail(400, { error: file.error });
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

            /**
             * Delete old cover
             */
            if (coverId) {
                const resDel = await upsendApi({
                    url: `/images/${coverId}`,
                    method: "DELETE",
                });
                if (!resDel.success) {
                    return fail(400, { error: resDel.error });
                }
            }

            /**
             * Upload new cover
             * @type {import("$lib/server/safe").Safe<import("$lib/types").UpsendImage>}
             */
            const file = await upsendApi({
                url: "/images",
                method: "POST",
                file: cover,
            });
            if (!file.success) {
                return fail(400, { error: file.error });
            }

            coverId = file.data.id;
            coverUrl = file.data.url;
        }

        /** @type {import('$lib/proto/proto/Profile').Profile} */
        const data = {
            id: getFormValue(form, "id"),
            username: getFormValue(form, "username"),
            about: getFormValue(form, "about"),
            resume_id: resumeId,
            cover_id: coverId,
            cover_url: coverUrl,
        };

        /** @type {import("$lib/server/safe").GrpcSafe<import("$lib/proto/proto/Profile").Profile__Output>} */
        const res = await new Promise((r) => {
            server.CreateProfile(
                data,
                createMetadata(locals.user.id),
                grpcSafe(r),
            );
        });

        if (!res.success) {
            if (res.fields) {
                return fail(400, { fields: res.fields });
            }
            return fail(400, { error: res.error });
        }

        /**
         * Send email with the data to the user
         * We don't check for errors cos the upsend API is not critical
         */
        await safe(
            upsendApi({
                url: "/emails",
                method: "POST",
                email: {
                    email_to: locals.user.email,
                    email_name: res.data.username,
                    email_subject: "You've updated your profile",
                    email_html: `
                <p>Hi ${res.data.username},</p>
                <p>You've updated your profile. You can view it <a href="https://sgsg.bearbyte.org/profile">here</a>.</p>
                <p>Thanks!</p>
                `,
                },
            }),
        );

        end();
        return {
            profile: res.data,
        };
    },
};
