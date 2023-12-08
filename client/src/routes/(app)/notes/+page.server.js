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
    const notesStream = server.GetNotesByUserId({}, metadata);

    /** @type {Promise<void>} */
    const p = new Promise((res, rej) => {
        notesStream.on("data", (data) => notes.push(data));
        notesStream.on("error", (err) => rej(err));
        notesStream.on("end", () => res());
    });
    const r = await safe(p);

    if (r.error) {
        return {
            error: r.msg,
            notes: [],
        };
    }
    return {
        notes: notes,
    };
}

/** @type {import('./$types').Actions} */
export const actions = {
    insert: async ({ locals, request }) => {
        const form = await request.formData();

        /** @type {import("$lib/proto/proto/Note").Note} */
        const data = {
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
};
