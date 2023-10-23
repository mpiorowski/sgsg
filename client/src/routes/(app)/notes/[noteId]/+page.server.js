import { error, fail } from "@sveltejs/kit";
import { getFormValue } from "$lib/helpers";
import { createMetadata } from "$lib/server/metadata";
import { server } from "$lib/server/grpc";
import { grpcSafe } from "$lib/safe";

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

/** @type {import('./$types').Actions}*/
export const actions = {
    update: async ({ locals, request }) => {
        const form = await request.formData();

        /** @type {import("$lib/proto/proto/Note").Note} */
        const data = {
            id: getFormValue(form, "id"),
            title: getFormValue(form, "title"),
            content: getFormValue(form, "content"),
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

        // throw redirect(302, "/notes");
        return {
            success: true,
        };
    },
};
