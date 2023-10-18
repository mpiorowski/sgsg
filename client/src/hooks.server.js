import { COOKIE_DOMAIN } from "$env/static/private";
import { grpcSafe } from "$lib/safe";
import { server } from "$lib/server/grpc";
import { logger, perf } from "$lib/server/logger";
import { createMetadata } from "$lib/server/metadata";
import { redirect } from "@sveltejs/kit";

/** @type {import('@sveltejs/kit').Handle} */
export async function handle({ event, resolve }) {
    const end = perf("Auth");
    event.locals.user = {
        id: "",
        created: "",
        updated: "",
        deleted: "",
        email: "",
        avatar: "",
        role: "",
        sub: "",
        subscriptionId: "",
        subscriptionEnd: "",
        _deleted: "deleted",
        _subscriptionEnd: "subscriptionEnd",
    };

    if (event.url.pathname === "/auth") {
        event.cookies.set("token", "", {
            domain: COOKIE_DOMAIN,
            path: "/",
            maxAge: 0,
        });
        return await resolve(event);
    }

    const token = event.cookies.get("token");
    if (!token) {
        logger.info("No token");
        throw redirect(302, "/auth");
    }

    const metadata = createMetadata(token);
    /** @type {import("$lib/safe").Safe<import("$lib/proto/proto/AuthResponse").AuthResponse__Output>} */
    const auth = await new Promise((res) => {
        server.Auth({}, metadata, grpcSafe(res));
    });
    if (auth.error || !auth.data.tokenId || !auth.data.user) {
        logger.error("Error during auth");
        throw redirect(302, "/auth");
    }

    event.locals.user = auth.data.user;
    event.locals.token = auth.data.tokenId;

    end();
    const response = await resolve(event);
    // max age is 30 days
    response.headers.append(
        "set-cookie",
        `token=${auth.data.tokenId}; HttpOnly; SameSite=Lax; Secure; Max-Age=2592000; Path=/; Domain=${COOKIE_DOMAIN}`,
    );
    return response;
}
