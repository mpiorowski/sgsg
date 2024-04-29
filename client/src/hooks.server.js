import { Role } from "$lib/proto/proto/Role";
import { authService, grpcSafe } from "$lib/server/grpc";
import { logger, perf } from "$lib/server/logger";
import { createMetadata } from "$lib/server/metadata";
import { redirect } from "@sveltejs/kit";
import { building } from "$app/environment";

/** @type {import('@sveltejs/kit').Handle} */
export async function handle({ event, resolve }) {
    logger.info(`path: ${event.url.pathname}`);
    if (building) {
        return await resolve(event);
    }

    const end = perf("auth");
    event.locals.user = {
        id: "",
        created: "",
        updated: "",
        deleted: "",
        email: "",
        avatar: "",
        role: Role.ROLE_UNSET,
        sub: "",
        subscription_id: "",
        subscription_end: "",
        subscription_check: "",
        subscription_active: false,
    };

    if (event.url.pathname === "/auth") {
        event.cookies.set("token", "", {
            path: "/",
            maxAge: 0,
        });
        return await resolve(event);
    }

    /**
     * Check if we have a token
     * If we don't have a token, redirect to the auth page
     */
    const token =
        event.url.searchParams.get("token") ?? event.cookies.get("token");
    if (!token) {
        logger.info("No token");
        throw redirect(302, "/auth");
    }

    const metadata = createMetadata(token);
    /** @type {import("$lib/server/safe").Safe<import("$lib/proto/proto/AuthResponse").AuthResponse__Output>} */
    const auth = await new Promise((res) => {
        authService.Auth({}, metadata, grpcSafe(res));
    });
    if (!auth.success) {
        logger.error(`Error during auth: ${auth.error}`);
        throw redirect(302, "/auth?error=unauthorized");
    }
    if (!auth.data.token || !auth.data.user) {
        logger.error("Error during auth");
        throw redirect(302, "/auth?error=unauthorized");
    }

    event.locals.user = auth.data.user;
    event.locals.token = auth.data.token;

    /** Last check to make sure we have a user */
    if (!event.locals.user.id) {
        logger.error("No user found");
        throw redirect(302, "/auth?error=unauthorized");
    }

    end();
    event.cookies.set("token", auth.data.token, {
        path: "/",
        // 7 days
        maxAge: 604800,
        sameSite: "lax",
        secure: true,
        httpOnly: true,
    });
    if (event.url.pathname === "/") {
        throw redirect(303, "/profile");
    }
    return await resolve(event, {
        preload: ({ type }) =>
            type === "js" || type === "css" || type === "font",
    });
}
