import { error } from "@sveltejs/kit";
import type { PageServerLoad, Actions } from "./$types";
import { fetchToken, notesClient, type Note, type UserId } from "src/grpc";
import type { NoteId } from "../../../../grpc/grpc/NoteId";
import type { Note__Output } from "../../../../grpc/grpc/Note";
import { URI_NOTES } from "$env/static/private";

export const load = (async ({ locals }) => {
    try {
        const request: UserId = { userId: locals.userId };
        const metadata = await fetchToken(URI_NOTES);
        const stream = notesClient.getNotes(request, metadata);
        const notes: Note__Output[] = [];

        const promise = new Promise<Note__Output[]>((resolve, reject) => {
            stream.on("data", (note) => {
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
    create: async ({ locals, request }) => {
        const form = await request.formData();
        const title = form.get("title");
        const content = form.get("content");

        if (!title || !content) {
            throw error(400, "Missing title or content");
        }

        try {
            const note: Note = {
                title: title as string,
                content: content as string,
                userId: locals.userId,
            };

            const metadata = await fetchToken(URI_NOTES);
            const promise = new Promise<Note__Output>((resolve, reject) => {
                notesClient.createNote(note, metadata, (err, response) => ((err || !response) ? reject(err) : resolve(response)));
            });

            return {
                note: await promise,
            };
        } catch (err) {
            console.error(err);
            throw error(500, "Could not create note");
        }
    },
    delete: async ({ locals, request }) => {
        const form = await request.formData();
        const id = form.get("id");

        if (!id) {
            throw error(400, "Missing id");
        }

        try {
            const request: NoteId = { noteId: id as string, userId: locals.userId };

            const metadata = await fetchToken(URI_NOTES);
            const promise = new Promise<Note__Output>((resolve, reject) => {
                notesClient.deleteNote(request, metadata, (err, response) => ((err || !response) ? reject(err) : resolve(response)));
            });

            return {
                note: await promise,
            };
        } catch (err) {
            console.error(err);
            throw error(500, "Failed to delete note");
        }
    },
} satisfies Actions;
