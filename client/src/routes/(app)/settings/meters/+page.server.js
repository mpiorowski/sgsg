import { fail } from "@sveltejs/kit";
import { createMetadata } from "$lib/server/metadata";
import { server } from "$lib/server/grpc";
import { grpcSafe } from "$lib/safe";
import { getFormValue } from "$lib/helpers";

/** @type {import('./$types').Actions}*/
export const actions = {
    createMeter: async ({ locals, request }) => {
        const form = await request.formData();
        /** @type {import("$lib/proto/proto/Meter").Meter} */
        const data = {
            id: getFormValue(form.get("id")),
            name: getFormValue(form.get("name")),
            description: getFormValue(form.get("description")),
        };

        const metadata = createMetadata(locals.token);
        /** @type {import("$lib/safe").Safe<import("$lib/proto/proto/Meter").Meter>} */
        const meter = await new Promise((res) =>
            server.CreateMeter(data, metadata, grpcSafe(res)),
        );

        if (meter.error) {
            if (meter.fields) {
                const errors = meter.fields.map((field) => ({
                    field: field.field + "_" + data.id,
                    tag: field.tag,
                }));
                return fail(400, { fields: errors });
            }
            return fail(400, { error: meter.msg });
        }

        return { meter: meter.data };
    },
    deleteMeter: async ({ locals, request }) => {
        const form = await request.formData();
        /** @type {import("$lib/proto/proto/Id").Id} */
        const data = { id: getFormValue(form.get("id")) };

        const metadata = createMetadata(locals.token);
        /** @type {import("$lib/safe").Safe<import("$lib/proto/proto/Id").Id>} */
        const meter = await new Promise((res) =>
            server.DeleteMeter(data, metadata, grpcSafe(res)),
        );

        if (meter.error) {
            return fail(400, { error: meter.msg });
        }

        return { meter: meter.data };
    },
};
