import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load = (({ locals }) => {
    if (!locals.user || !locals.user.id) {
        throw redirect(303, "/login");
    }
    return { user: locals.user };
}) satisfies PageServerLoad;
