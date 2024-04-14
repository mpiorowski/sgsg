import { getAllValues, getValue } from "$lib/helpers";
import { grpcSafe, profileService } from "$lib/server/grpc";
import { perf } from "$lib/server/logger";
import { createMetadata } from "$lib/server/metadata";
import { error, fail } from "@sveltejs/kit";

/** @type {import('./$types').PageServerLoad} */
export async function load({ locals }) {
    const end = perf("load_profile");
    const metadata = createMetadata(locals.user.id);
    /** @type {import("$lib/server/safe").GrpcSafe<import("$lib/proto/proto/Profile").Profile__Output>} */
    const profile = await new Promise((r) =>
        profileService.GetProfile({}, metadata, grpcSafe(r)),
    );
    if (!profile.success) {
        throw error(500, profile.error);
    }
    end();
    return {
        profile: profile.data,
    };
}

/** @type {import('./$types').Actions} */
export const actions = {
    update_profile: async ({ request, locals }) => {
        const end = perf("update_profile");
        const form = await request.formData();

        /** @type {import("$lib/proto/proto/Profile").Profile} */
        const data = {
            id: getValue(form, "id"),
            active: getValue(form, "active") === "on",
            username: getValue(form, "username"),
            about: getValue(form, "about"),
            first_name: getValue(form, "first_name"),
            last_name: getValue(form, "last_name"),
            email: getValue(form, "email"),
            country: getValue(form, "country"),
            street_address: getValue(form, "street_address"),
            city: getValue(form, "city"),
            state: getValue(form, "state"),
            zip: getValue(form, "zip"),
            email_notifications: getAllValues(form, "email_notifications").join(
                ",",
            ),
            push_notification: getValue(form, "push_notification"),
            resume: "",
            cover: "",
            position: getValue(form, "position"),
            skills: getValue(form, "skills"),
        };

        const metadata = createMetadata(locals.user.id);
        /** @type {import("$lib/server/safe").GrpcSafe<import("$lib/proto/proto/Profile").Profile__Output>} */
        const profile = await new Promise((r) =>
            profileService.UpdateProfile(data, metadata, grpcSafe(r)),
        );
        if (!profile.success) {
            return fail(500, { error: profile.error, fields: profile.fields });
        }
        end();
        return { profile: profile.data };
    },
};
