import { apiRequest } from "$lib/api.util";
import { error } from "@sveltejs/kit";
import type { PageServerLoad, Actions } from "./$types";

import grpc from "@grpc/grpc-js";
import protoLoader from "@grpc/proto-loader";
import type { ProtoGrpcType } from "../../../../server/grpc/grpc";
import type { UserId } from "../../../../server/grpc/grpc/UserId";
import type { Note } from "../../../../server/grpc/grpc/Note";

const host = 'service-notes';
const packageDefinition = protoLoader.loadSync('../../../../server/grpc/grpc.proto');
const proto = grpc.loadPackageDefinition(
  packageDefinition
) as unknown as ProtoGrpcType;

const client = new proto.grpc.NotesService(
  host,
  grpc.credentials.createInsecure()
);
export const load = (async ({ cookies }) => {
    try {
        const request: UserId = { userId: "test" };
        const stream = client.getNotes(request);
        const notes: Note[] = [];

        const promise = new Promise((resolve, reject) => {
            stream.on("data", (note: Note) => {
                notes.push(note);
            });

            stream.on("end", () => {
                resolve(notes);
            });

            stream.on("error", (err: unknown) => {
                reject(err);
            });
        });

        const data = await promise;

        return {
            notes: data,
        };
    } catch (err) {
        console.error(err);
        throw error(500, "Could not load notes");
    }
}) satisfies PageServerLoad;

export const actions = {
    create: async ({ cookies, request }) => {
        const form = await request.formData();
        const title = form.get("title");
        const content = form.get("content");

        if (!title || !content) {
            throw error(400, "Missing title or content");
        }

        try {
            const response = await apiRequest<Note>({
                url: "/notes",
                method: "POST",
                body: JSON.stringify({ title, content }),
                cookies,
            });
            return {
                note: response,
            };
        } catch (err) {
            console.error(err);
            throw error(500, "Could not create note");
        }
    },
    delete: async ({ cookies, request }) => {
        const form = await request.formData();
        const id = form.get("id");

        if (!id) {
            throw error(400, "Missing id");
        }

        try {
            const response = await apiRequest<Note>({
                url: "/notes/" + id,
                method: "DELETE",
                cookies,
            });
            return {
                note: response,
            };
        } catch (err) {
            console.error(err);
            throw error(500, "Failed to delete note");
        }
    },
} satisfies Actions;
