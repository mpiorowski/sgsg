import type { HandleServerError } from "@sveltejs/kit";
import { SvelteKitAuth } from "@auth/sveltekit";
import { GOOGLE_ID, GOOGLE_SECRET } from "$env/static/private";
import { redirect, type Handle } from "@sveltejs/kit";
import { sequence } from "@sveltejs/kit/hooks";
import Google from "@auth/core/providers/google";
import type { Provider } from "@auth/core/providers";
import { usersClient } from "./grpc";
import type { AuthRequest } from "../../grpc/grpc/AuthRequest";

// TODO - finish handle error
export const handleError: HandleServerError = ({ error }) => {
    console.error(error);
    return {
        message: "Whoops!",
        code: "UNKNOWN",
    };
};

export const authorization = (async ({ event, resolve }) => {
    try {
        const session = await event.locals.getSession();
        if (!session?.user?.email) {
            throw new Error("No user session");
        }
        const request: AuthRequest = {
            providerId: "not using this",
            email: session.user.email,
        };

        const promise = new Promise<void>((resolve, reject) => {
            usersClient.Auth(request, (err, response) => {
                if (err || !response?.id || !response?.role) {
                    return reject(err);
                }
                event.locals.userId = response.id;
                event.locals.role = response.role;
                resolve();
            });
        });
        await promise;
    } catch (err) {
        console.error("User is not authorized: %s", err);
        event.locals.userId = "";
        event.locals.role = "";
        if (!event.url.pathname.startsWith("/auth")) {
            throw redirect(303, "/auth");
        }
    }

    if (event.url.pathname.startsWith("/auth") && event.locals.userId) {
        throw redirect(303, "/");
    }

    // If the request is still here, just proceed as normally
    const result = await resolve(event, {
        transformPageChunk: ({ html }) => html,
    });
    return result;
}) satisfies Handle;

// First handle authentication, then authorization
// Each function acts as a middleware, receiving the request handle
// And returning a handle which gets passed to the next function
export const handle: Handle = sequence(
    SvelteKitAuth({
        providers: [
            Google({
                clientId: GOOGLE_ID,
                clientSecret: GOOGLE_SECRET,
            }) as Provider,
        ],
    }),
    authorization,
);
