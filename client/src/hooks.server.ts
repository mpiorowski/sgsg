import { apiRequest } from "$lib/api.util";
import type { Handle, HandleServerError } from "@sveltejs/kit";
import { Config } from "./config";
import type { User } from "./types/user.type";

// TODO - finish handle error
export const handleError: HandleServerError = ({ error, event }) => {
    console.error(error);
    return {
        message: 'Whoops!',
        code: 'UNKNOWN'
    };
};

export const handle: Handle = async ({ event, resolve }) => {
    try {
        const session = event.cookies.get("sessionCookie");
        if (!session) {
            console.info("No session cookie");
            event.locals.user = null;
            return await resolve(event);
        }
        const user = await apiRequest<User>({
            url: "/auth",
            method: "GET",
            cookies: event.cookies,
        });
        event.locals.user = {
            id: user.id,
            email: user.email,
            role: user.role,
            providerId: user.providerId,
        }
    } catch (error) {
        console.error(error);
        event.locals.user = null;
    }
    return await resolve(event);
};
