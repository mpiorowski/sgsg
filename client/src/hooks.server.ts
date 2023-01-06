import type { HandleServerError } from "@sveltejs/kit";
import { SvelteKitAuth } from "@auth/sveltekit";
import { GOOGLE_ID, GOOGLE_SECRET } from "$env/static/private";
import { redirect, type Handle } from "@sveltejs/kit";
import { sequence } from "@sveltejs/kit/hooks";
import Google from "@auth/core/providers/google"
import type { Provider } from "@auth/core/providers";

// TODO - finish handle error
export const handleError: HandleServerError = ({ error }) => {
    console.error(error);
    return {
        message: "Whoops!",
        code: "UNKNOWN",
    };
};

export const authorization = (async ({ event, resolve }) => {
    console.info("Checking authorization");
    // Protect any routes under /authenticated
    if (!event.url.pathname.startsWith("/auth")) {
        const session = await event.locals.getSession();
        if (!session) {
            throw redirect(303, "/auth");
        }
    }

    // If the request is still here, just proceed as normally
    const result = await resolve(event, {
        transformPageChunk: ({ html }) => html,
    });
    return result;
}) satisfies Handle

// First handle authentication, then authorization
// Each function acts as a middleware, receiving the request handle
// And returning a handle which gets passed to the next function
export const handle: Handle = sequence(
    SvelteKitAuth({
        providers: [
            Google({ clientId: GOOGLE_ID, clientSecret: GOOGLE_SECRET }) as Provider,
        ],
    }),
    authorization,
);
