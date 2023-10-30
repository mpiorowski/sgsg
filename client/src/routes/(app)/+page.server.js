import { getFormValue } from "$lib/helpers";
import { grpcSafe } from "$lib/safe";
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

        const resume = form.get("resume");
        if (!(resume instanceof File) || resume.size <= 0) {
            return fail(400, { error: "Resume is required" });
        }
        if (resume.size > 5 * 1024 * 1024) {
            return fail(400, { error: "Resume must be less than 5MB" });
        }
        if (!resume.name.endsWith(".pdf")) {
            return fail(400, { error: "Resume must be a PDF" });
        }

        const cover = form.get("cover");
        if (!(cover instanceof File) || cover.size <= 0) {
            return fail(400, { error: "Cover is required" });
        }
        if (cover.size > 5 * 1024 * 1024) {
            return fail(400, { error: "Cover must be less than 5MB" });
        }
        const extensions = [".png", ".jpg", ".jpeg", ".gif", ".svg"];
        if (!extensions.some((ext) => cover.name.endsWith(ext))) {
            return fail(400, { error: "Cover must be an image" });
        }

        /** @type {import('$lib/proto/proto/Profile').Profile} */
        const data = {
            id: getFormValue(form, "id"),
            username: getFormValue(form, "username"),
            about: getFormValue(form, "about"),
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
            return fail(400, { error: res.msg });
        }

        return {
            profile: res.data,
        };
    },
};
