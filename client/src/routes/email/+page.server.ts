import type { Actions } from "./$types";
import { NODE_ENV } from "$env/static/private";
import { pubSubClient } from "src/pubsub";
import { error } from "@sveltejs/kit";


export const actions = {
    sendEmail: async ({ request }) => {
        if (NODE_ENV === "development") {
            return {
                message: "Email sent",
            };
        }

        const form = await request.formData();
        const email = form.get("email") as string;
        const message = form.get("message") as string;

        const data = {
            to: email,
            type: "CONTACT",
            html: [message]
        };
        const dataBuffer = Buffer.from(JSON.stringify(data));

        try {
            const messageId = await pubSubClient
                .topic("email")
                .publishMessage({ data: dataBuffer });
            console.info('Message published: %s', messageId);
        } catch (err) {
            console.error(err);
            throw error(409, "Error sending email");
        }
    }
} satisfies Actions;
