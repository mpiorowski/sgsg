import { safe } from "$lib/safe";
import { server } from "$lib/server/grpc";
import { createMetadata } from "$lib/server/metadata";

/** @type {import('./$types').LayoutServerLoad} */
export async function load({ locals }) {
    const stream = server.GetMeters({}, createMetadata(locals.token));
    /** @type {import("$lib/proto/proto/Meter").Meter__Output[]} */
    const meters = [];
    const grpc = await safe(
        /** @type {Promise<void>} */ (
            new Promise((res, rej) => {
                stream.on("data", (data) => meters.push(data));
                stream.on("error", (err) => rej(err));
                stream.on("end", () => res());
            })
        ),
    );

    if (grpc.error) {
        return { meters: [], error: grpc.msg };
    }

    return {
        meters: meters,
    };
}
