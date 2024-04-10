import { authService, grpcSafe } from "$lib/server/grpc";
import { perf } from "$lib/server/logger";
import { createMetadata } from "$lib/server/metadata";
import { fail, redirect } from "@sveltejs/kit";

/** @type {import('./$types').PageServerLoad} */
export function load({ locals }) {
    return {
        subscription_active: locals.user.subscription_active,
    };
}

/** @type {import('./$types').Actions} */
export const actions = {
    createStripeCheckout: async ({ locals }) => {
        const end = perf("create_stripe_checkout");
        const metadata = createMetadata(locals.token);

        /** @type {import("$lib/server/safe").Safe<import("$lib/proto/proto/StripeUrlResponse").StripeUrlResponse__Output>} */
        const s = await new Promise((r) =>
            authService.CreateStripeCheckout({}, metadata, grpcSafe(r)),
        );

        if (!s.success) {
            return fail(500, { error: s.error });
        }

        end();
        throw redirect(303, s.data.url ?? "");
    },
    createStripePortal: async ({ locals }) => {
        const end = perf("create_stripe_portal");
        const metadata = createMetadata(locals.token);

        /** @type {import("$lib/server/safe").Safe<import("$lib/proto/proto/StripeUrlResponse").StripeUrlResponse__Output>} */
        const s = await new Promise((r) =>
            authService.CreateStripePortal({}, metadata, grpcSafe(r)),
        );

        if (!s.success) {
            return fail(500, { error: s.error });
        }

        end();
        throw redirect(303, s.data.url ?? "");
    },
};
