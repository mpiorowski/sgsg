import { getFormValue } from "$lib/helpers";
import { safe } from "$lib/server/safe";
import { grpcSafe, server } from "$lib/server/grpc";
import { createMetadata } from "$lib/server/metadata";
import { error, fail } from "@sveltejs/kit";
import { perf } from "$lib/server/logger";

/** @type {import('./$types').PageServerLoad} */
export async function load({ locals }) {
    const end = perf("load_notes");
    /** @type {import("$lib/proto/proto/Note").Note__Output[]} */
    const notes = [];
    const metadata = createMetadata(locals.user.id);
    const notesStream = server.GetNotesByUserId({}, metadata);

    /** @type {Promise<void>} */
    const p = new Promise((res, rej) => {
        notesStream.on("data", (data) => notes.push(data));
        notesStream.on("error", (err) => rej(err));
        notesStream.on("end", () => res());
    });
    const r = await safe(p);

    if (!r.success) {
        throw error(500, r.error);
    }
    end();
    return {
        notes: notes,
    };
}

/** @type {import('./$types').Actions} */
export const actions = {
    insert: async ({ locals, request }) => {
        const end = perf("insert_note");
        const form = await request.formData();

        /** @type {import("$lib/proto/proto/Note").Note} */
        const data = {
            title: getFormValue(form, "title"),
            content: getFormValue(form, "content"),
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
};
