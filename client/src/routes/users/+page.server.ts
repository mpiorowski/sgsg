import { URI_USERS } from "$env/static/private";
import { error, redirect } from "@sveltejs/kit";
import { fetchToken, usersClient } from "src/grpc";
import type { User__Output } from "../../../../proto/proto/User";
import { UserRole } from "../../../../proto/proto/UserRole";
import type { PageServerLoad } from "./$types";

export const load = (async ({ locals }) => {
    if (locals.role !== UserRole.ROLE_ADMIN) {
        throw redirect(303, "/");
    }
    try {
        const metadata = await fetchToken(URI_USERS);
        const stream = usersClient.GetUsers(metadata);
        const users: User__Output[] = [];

        const promise = new Promise<User__Output[]>((resolve, reject) => {
            stream.on("data", (user: User__Output) => {
                users.push(user);
            });
            stream.on("end", () => {
                resolve(users);
            });
            stream.on("error", (err) => {
                reject(err);
            });
        });
        return {
            users: await promise,
        };
    } catch (err) {
        console.error(err);
        throw error(500, "Could not load users");
    }
}) satisfies PageServerLoad;
