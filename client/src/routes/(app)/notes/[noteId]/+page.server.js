import { error, fail } from "@sveltejs/kit";
import { getFormValue } from "$lib/helpers";
import { createMetadata } from "$lib/server/metadata";
import { grpcSafe, server } from "$lib/server/grpc";
import { schema } from "./note";
import { perf } from "$lib/server/logger";

/** @type {import('./$types').PageServerLoad} */
export async function load({ locals, params }) {
    const end = perf("load_note");
    const id = params.noteId;
    if (!id) {
        throw error(409, "Missing note id");
    }

    const metadata = createMetadata(locals.user.id);
    /** @type {import("$lib/server/safe").GrpcSafe<import("$lib/proto/proto/Note").Note__Output>} */
    const req = await new Promise((r) => {
        server.GetNoteById({ id }, metadata, grpcSafe(r));
    });

    if (!req.success) {
        throw error(404, req.error);
    }

    end();
    return {
        note: req.data,
    };
}

/** @type {import('./$types').Actions} */
export const actions = {
    update: async ({ locals, request }) => {
        const end = perf("update_note");
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

        const metadata = createMetadata(locals.user.id);
        /** @type {import("$lib/server/safe").GrpcSafe<import("$lib/proto/proto/Note").Note__Output>} */
        const req = await new Promise((r) => {
            server.CreateNote(data, metadata, grpcSafe(r));
        });

        if (!req.success) {
            return fail(500, { error: req.error, fields: req.fields });
        }
        end();
        return { note: req.data };
    },
    delete: async ({ locals, request }) => {
        const end = perf("delete_note");
        const form = await request.formData();
        /** @type {import("$lib/proto/proto/Id").Id} */
        const data = {
            id: getFormValue(form, "id"),
        };
        const metadata = createMetadata(locals.user.id);
        /** @type {import("$lib/server/safe").GrpcSafe<import("$lib/proto/proto/Empty").Empty__Output>} */
        const req = await new Promise((r) => {
            server.DeleteNoteById(data, metadata, grpcSafe(r));
        });

        if (!req.success) {
            return fail(400, { error: req.error });
        }

        end();
        return {
            success: true,
        };
    },
};
