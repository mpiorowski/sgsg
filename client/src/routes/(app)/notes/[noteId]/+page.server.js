import { error, fail } from "@sveltejs/kit";
import { getFormValue } from "$lib/helpers";
import { createMetadata } from "$lib/server/metadata";
import { server } from "$lib/server/grpc";
import { grpcSafe } from "$lib/safe";
import { schema } from "./note";

/** @type {import('./$types').PageServerLoad} */
export async function load({ locals, params }) {
    const id = params.noteId;
    if (!id) {
        throw error(409, "Missing note id");
    }

    const metadata = createMetadata(locals.token);
    /** @type {import("$lib/safe").Safe<import("$lib/proto/proto/Note").Note__Output>} */
    const req = await new Promise((r) => {
        server.GetNoteById({ id }, metadata, grpcSafe(r));
    });

    if (req.error) {
        throw error(404, req.msg);
    }

    return {
        note: req.data,
    };
}

/** @type {import('./$types').Actions} */
export const actions = {
    update: async ({ locals, request }) => {
        const form = await request.formData();

        const validation = schema.safeParse({
            id: form.get("id"),
            title: form.get("title"),
            content: form.get("content"),
        });

        if (!validation.success) {
            const fieldErrors = validation.error.flatten().fieldErrors;
            const fields = Object.entries(fieldErrors).map(([key, value]) => ({
                field: key,
                tag: value.join(", "),
            }));
            return fail(400, { fields: fields });
        }

        /** @type {import("./type").Note} */
        const data = {
            id: validation.data.id,
            title: validation.data.title,
            content: validation.data.content,
        };

        const metadata = createMetadata(locals.token);
        /** @type {import("$lib/safe").Safe<import("$lib/proto/proto/Note").Note__Output>} */
        const req = await new Promise((r) => {
            server.CreateNote(data, metadata, grpcSafe(r));
        });

        if (req.error) {
            if (req.fields) {
                return fail(400, { fields: req.fields });
            }
            return fail(500, { error: req.msg });
        }

        return { note: req.data };
    },
    delete: async ({ locals, request }) => {
        const form = await request.formData();
        /** @type {import("$lib/proto/proto/Id").Id} */
        const data = {
            id: getFormValue(form, "id"),
        };
        const metadata = createMetadata(locals.token);
        /** @type {import("$lib/safe").Safe<import("$lib/proto/proto/Empty").Empty>} */
        const req = await new Promise((r) => {
            server.DeleteNote(data, metadata, grpcSafe(r));
        });

        if (req.error) {
            return fail(400, { error: req.msg });
        }

        return {
            success: true,
        };
    },
};
