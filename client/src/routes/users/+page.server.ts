import { apiRequest } from "$lib/api.util";
import { error, redirect } from "@sveltejs/kit";
import type { User } from "src/types/user.type";
import type { PageServerLoad, Actions } from "./$types";

export const load = (async ({ locals, cookies }) => {
    if (!locals.user?.id) {
        throw redirect(303, "/login");
    }

    try {
        const users = await apiRequest<User[]>({
            url: "/users",
            method: "GET",
            cookies
        });
        return { users: users };
    }
    catch (err) {
        console.error(err);
        throw error(500, "Could not load users");
    }
}) satisfies PageServerLoad;

export const actions = {
    delete: async ({ locals, cookies, request }) => {
        if (!locals.user?.id) {
            throw redirect(303, "/login");
        }

        const form = await request.formData();
        const user = form.get("user");

        try {
            const response = await apiRequest<User>({
                url: "/users",
                method: "DELETE",
                body: user as string,
                cookies,
            });
            return {
                user: response,
            }
        } catch (err) {
            console.error(err);
            throw error(500, "Failed to delete user");
        }
    }
} satisfies Actions;
