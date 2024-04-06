import { getValue } from "$lib/helpers";
import { safe } from "$lib/server/safe";
import { grpcSafe, profileService } from "$lib/server/grpc";
import { createMetadata } from "$lib/server/metadata";
import { error, fail } from "@sveltejs/kit";
import { perf } from "$lib/server/logger";

/** @type {import('./$types').PageServerLoad} */
export async function load({ locals, url }) {
    const end = perf("load_notes");
    const page = url.searchParams.get("page") || "1";
    const limit = url.searchParams.get("limit") || "2";

    const total = await new Promise((res) =>
        profileService.CountNotesByUserId(
            {},
            createMetadata(locals.token),
            grpcSafe(res),
        ),
    );
    if (!total.success) {
        throw error(500, total.error);
    }

    /** @type {import("$lib/proto/proto/Page").Page} */
    const request = {
        offset: (parseInt(page) - 1) * parseInt(limit),
        limit: parseInt(limit),
    };
    /** @type {import("$lib/proto/proto/Note").Note__Output[]} */
    const notes = [];
    const metadata = createMetadata(locals.token);
    const notesStream = profileService.GetNotesByUserId(request, metadata);

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
        total: total.data.count,
        limit: parseInt(limit),
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
            title: getValue(form, "title"),
            content: getValue(form, "content"),
        };
        const metadata = createMetadata(locals.token);
        /** @type {import("$lib/server/safe").GrpcSafe<import("$lib/proto/proto/Note").Note__Output>} */
        const req = await new Promise((r) => {
            profileService.CreateNote(data, metadata, grpcSafe(r));
        });

        if (!req.success) {
            return fail(500, { error: req.error, fields: req.fields });
        }

        end();
        return { note: req.data };
    },
};
