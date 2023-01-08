import { URI_FILES } from "$env/static/private";
import { error, type Actions } from "@sveltejs/kit";
import { fetchToken, filesClient } from "src/grpc";
import type { PageServerLoad } from "../$types";
import type { File__Output, File } from "../../../../proto/proto/File";
import type { TargetId } from "../../../../proto/proto/TargetId";

export const load = (async ({ locals }) => {
    const { userId } = locals;

    const data: TargetId = {
        targetId: userId,
    };

    const files: File__Output[] = [];
    const metadata = await fetchToken(URI_FILES);
    const stream = filesClient.getFiles(data, metadata);
    const promise = new Promise<File__Output[]>((resolve, reject) => {
        stream.on("data", (file) => {
            files.push(file);
        });

        stream.on("end", () => {
            resolve(files);
        });

        stream.on("error", (err: unknown) => {
            reject(err);
        });
    });
    return {
        files: await promise,
    };
}) satisfies PageServerLoad;

export const actions = {
    file: async ({ request, locals }) => {
        try {
            const form = await request.formData();
            const file = form.get("file") as Blob;
            const buffer = Buffer.from(await file.arrayBuffer());

            const data: File = {
                name: file.name,
                type: "USER",
                targetId: locals.userId,
                data: buffer,
            };

            const metadata = await fetchToken(URI_FILES);
            const promise = new Promise<File__Output>((resolve, reject) => {
                filesClient.CreateFile(data, metadata, (err, res) =>
                    err || !res ? reject(err) : resolve(res),
                );
            });

            return {
                file: await promise,
            };
        } catch (err) {
            console.error(err);
            throw error(409, "Error creating file");
        }
    },
} satisfies Actions;
