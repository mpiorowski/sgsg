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
        /** @type {Promise<void>} */(
            new Promise((res, rej) => {
                notesStream.on("data", (data) => notes.push(data));
                notesStream.on("error", (err) => rej(err));
                notesStream.on("end", () => res());
            })
        ),
    );
    if (req.error) {
        return fail(400, {
            ...req,
            notes: [],
        });
    }
    return {
        note: req.data,
    };
}

/** @type {import('./$types').Actions} */
export const actions = {
    default: async ({ locals, request }) => {
        const form = await request.formData();

        /** @type {import("$lib/proto/proto/Note").Note} */
        const data = {
            title: getFormValue(form, "title"),
            content: getFormValue(form, "content"),
        };
        const metadata = createMetadata(locals.token);
        /** @type {import("$lib/safe").Safe<import("$lib/proto/proto/Note").Note>} */
        const req = await new Promise((r) => {
            server.CreateNote(data, metadata, grpcSafe(r));
        });

        if (req.error) {
            return fail(400, req);
        }

        return { note: req.data };
    },
};
