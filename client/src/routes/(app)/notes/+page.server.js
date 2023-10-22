import { getFormValue } from "$lib/helpers";
import { grpcSafe, safe } from "$lib/safe";
import { server } from "$lib/server/grpc";
import { createMetadata } from "$lib/server/metadata";
import { fail } from "@sveltejs/kit";

/** @type {import('./$types').PageServerLoad} */
export async function load({ locals }) {
    /** @type {import("$lib/proto/proto/Note").Note__Output[]} */
    const notes = [];
    const metadata = createMetadata(locals.token);
    const notesStream = server.GetNotes({}, metadata);
    const req = await safe(
        /** @type {Promise<void>} */ (
            new Promise((res, rej) => {
                notesStream.on("data", (data) => notes.push(data));
                notesStream.on("error", (err) => rej(err));
                notesStream.on("end", () => res());
            })
        ),
    );
    if (req.error) {
        return {
            error: req.msg,
            notes: [],
        };
    }
    return {
        notes: notes,
    };
}

/** @type {import('./$types').Actions} */
export const actions = {
    create: async ({ locals, request }) => {
        const form = await request.formData();

        /** @type {import("$lib/proto/proto/Note").Note} */
        const data = {
            id: getFormValue(form, "id"),
            title: getFormValue(form, "title"),
            content: getFormValue(form, "content"),
        };
        const metadata = createMetadata(locals.token);
        /** @type {import("$lib/safe").Safe<import("$lib/proto/proto/Note").Note>} */
        const req = await new Promise((r) => {
            server.CreateNote(data, metadata, grpcSafe(r));
        });

        if (req.error) {
            return fail(400, { error: req.msg });
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

        return { success: true };
    },
};
