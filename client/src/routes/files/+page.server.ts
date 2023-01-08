import { URI_FILES } from "$env/static/private";
import { error, type Actions } from "@sveltejs/kit";
import { fetchToken, filesClient } from "src/grpc";
import type { File } from "../../../../proto/proto/File";

export const actions = {
    file: async ({ request }) => {
        try {
            const form = await request.formData();
            const file = form.get("file") as Blob;
            const buffer = Buffer.from(await file.arrayBuffer());

            const data: File = {
                name: file.name,
                type: "USER",
                data: buffer,
            }

            const metadata = await fetchToken(URI_FILES);
            const promise = await new Promise((resolve, reject) => {
                filesClient.CreateFile(data, metadata, (err, res) => (err ? reject(err) : resolve(res)));
            });

            return {
                file: await promise,
            };
        } catch (err) {
            console.error(err);
            throw error(401, { message: "Error logging in" });
        }
    },
} satisfies Actions;
