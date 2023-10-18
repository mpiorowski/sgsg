import { fail } from "@sveltejs/kit";

/** @type {import('./$types').PageServerLoad} */
export function load({ locals }) {

    return {
        locals,
    };
}

/** @type {import('./$types').Actions}*/
export const actions = {
    default: async ({ request }) => {
        const form = await request.formData();

        if (!form.has("id")) {
            return fail(404, { error: "Not found" });
        }

        return { body: { id: form.get("id") } };
    },
};
