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
    delete: async ({ locals, cookies, params }) => {
        if (!locals.user?.id) {
            throw redirect(303, "/login");
        }

        const response = await fetch(Config.VITE_API_URL + "/users", {
            method: "DELETE",
            headers: {
                Cookie: `sessionCookie=${cookies.get("sessionCookie")}`,
            },
        });
        if (!response.ok) {
            console.error(await response.json());
            throw error(500, "Failed to delete user");
        }
        return redirect(303, "/users");
    }
} satisfies Actions;
