import type { Actions, PageServerLoad } from "./$types";
import { error, redirect } from "@sveltejs/kit";
import { apiRequest } from "$lib/api.util";

export const load = (({ locals }) => {
    if (locals.user?.id) {
        throw redirect(303, "/");
    }
    return { user: locals.user };
}) satisfies PageServerLoad;

export const actions = {
    login: async ({ request, cookies }) => {
        try {
            const form = await request.formData();
            const idToken = form.get("idToken");

            if (!idToken || typeof idToken !== "string") {
                throw new Error("idToken is not a string");
            }

            // TODO - add csrf protection
            // const csrfToken = form.get("csrfToken");
            // if (!csrfToken || typeof csrfToken !== 'string' || csrfToken !== cookies.get("csrfToken")) {
            // return { status: 403, message: "Invalid CSRF token" };
            // }

            const sessionCookie = await apiRequest<{ cookie: string }>({
                url: "/login",
                method: "POST",
                body: JSON.stringify({ idToken }),
            });

            // TODO - config cookie
            cookies.set("sessionCookie", sessionCookie.cookie, {
                httpOnly: true,
                secure: true,
                sameSite: "strict",
                maxAge: 60 * 60 * 24 * 5, // 5 days
            });
        } catch (err) {
            console.error(err);
            throw error(401, { message: "Error logging in" });
        }
        // TODO - this is not working on the client, cos not using form
        throw redirect(303, "/");
    },
    logout: async ({ cookies }) => {
        cookies.set("sessionCookie", "", {
            httpOnly: true,
            secure: true,
            sameSite: "strict",
            maxAge: 0,
        });
        throw redirect(303, "/login");
    },
} satisfies Actions;
