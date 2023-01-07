import type { LayoutServerLoad } from "./$types";

export const load = (async ({ locals }) => {
    return { session: await locals.getSession(), userId: locals.userId, role: locals.role };
}) satisfies LayoutServerLoad;
