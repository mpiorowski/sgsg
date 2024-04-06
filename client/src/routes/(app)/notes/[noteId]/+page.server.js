import { error, fail } from "@sveltejs/kit";
import { getValue } from "$lib/helpers";
import { createMetadata } from "$lib/server/metadata";
import { grpcSafe, profileService } from "$lib/server/grpc";
import { perf } from "$lib/server/logger";

/** @type {import('./$types').PageServerLoad} */
export async function load({ locals, params }) {
    const end = perf("load_note");
    const id = params.noteId;

    const metadata = createMetadata(locals.token);
    /** @type {import("$lib/server/safe").GrpcSafe<import("$lib/proto/proto/Note").Note__Output>} */
    const req = await new Promise((r) => {
        profileService.GetNoteById({ id }, metadata, grpcSafe(r));
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

        /** @type {import("$lib/proto/proto/Note").Note} */
        const data = {
            id: getValue(form, "id"),
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
    delete: async ({ locals, request }) => {
        const end = perf("delete_note");
        const form = await request.formData();
        /** @type {import("$lib/proto/proto/Id").Id} */
        const data = {
            id: getValue(form, "id"),
        };
        const metadata = createMetadata(locals.token);
        /** @type {import("$lib/server/safe").GrpcSafe<import("$lib/proto/proto/Empty").Empty__Output>} */
        const req = await new Promise((r) => {
            profileService.DeleteNoteById(data, metadata, grpcSafe(r));
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
