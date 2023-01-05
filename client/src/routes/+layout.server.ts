import type { LayoutServerLoad } from "./$types";

export const load = (async ({ locals }) => {
    return { session: await locals.getSession() };
}) satisfies LayoutServerLoad;
