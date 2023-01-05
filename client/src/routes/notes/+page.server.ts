import { apiRequest } from "$lib/api.util";
import { error, redirect } from "@sveltejs/kit";
import type { Note } from "src/types/note.type";
import type { PageServerLoad, Actions } from "./$types";

export const load = (async ({ cookies }) => {
    console.log(cookies);
    try {
        const notes = await apiRequest<Note[]>({
            url: "/notes",
            method: "GET",
            cookies
        });
        return { notes: notes || [] };
    }
    catch (err) {
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
            }
        } catch (err) {
            console.error(err);
            throw error(500, "Could not create note");
        }
    },
    delete: async ({ locals, cookies, request }) => {
        if (!locals.user?.id) {
            throw redirect(303, "/login");
        }

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
            }
        } catch (err) {
            console.error(err);
            throw error(500, "Failed to delete note");
        }
    }
} satisfies Actions;
