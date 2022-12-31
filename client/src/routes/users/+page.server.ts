import { apiRequest } from "$lib/api.util";
import { error, redirect } from "@sveltejs/kit";
import { Config } from "src/config";
import type { User } from "src/types/user.type";
import type { PageServerLoad, Actions } from "./$types";

export const load = (async ({ locals, cookies }) => {
    if (!locals.user?.id) {
        throw redirect(303, "/login");
    }

    const response = await fetch(Config.VITE_API_URL + "/users", {
        headers: {
            Cookie: `sessionCookie=${cookies.get("sessionCookie")}`,
        },
    });
    if (!response.ok) {
        console.error(await response.json());
        throw error(500, "Failed to get users");
    }
    const users = await response.json() as User[];
    return { users: users };
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
