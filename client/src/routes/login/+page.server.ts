import type { Actions, PageServerLoad } from "./$types";
import { error, redirect } from "@sveltejs/kit";
import { serverAuth } from "$lib/firebase.server";
import { Config } from "src/config";

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

            // Set session expiration to 5 days.
            const expiresIn = 60 * 60 * 24 * 5 * 1000;

            const sessionCookie = await serverAuth.createSessionCookie(
                idToken,
                { expiresIn },
            );

            // Add user to server
            const response = await fetch(Config.VITE_API_URL + "/auth", {
                method: "GET",
                headers: {
                    "Cookie": `sessionCookie=${sessionCookie}`,
                },
            });
            if (!response.ok) {
                throw new Error("Failed to add user to server");
            }

            // TODO - config cookie
            cookies.set("sessionCookie", sessionCookie, {
                httpOnly: true,
                secure: true,
                sameSite: "strict",
                maxAge: expiresIn,
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
