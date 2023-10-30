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
