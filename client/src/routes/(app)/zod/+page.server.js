import { z } from "zod";
import { fail } from "@sveltejs/kit";

/** @type {import('./$types').PageServerLoad} */
export function load({ locals }) {
    return {
        locals,
    };
}

/** @type {import('./$types').Actions}*/
export const actions = {
    default: async ({ locals, request }) => {
        const form = await request.formData();
        const id = form.get("id");

        const schema = z
            .object({
                id: z.string().uuid(),
            })
            .safeParse({ id });

        if (!schema.success) {
            return fail(400, { fields: schema.error.flatten().fieldErrors });
        }

        if (!locals.user) {
            return fail(401, { error: "Unauthorized" });
        }

        return { id };
    },
};
