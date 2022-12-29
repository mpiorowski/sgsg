import type { Handle } from "@sveltejs/kit";
import { credential, initializeApp, type ServiceAccount } from "firebase-admin";
import serviceAccount from "../../service.json";

export const handle: Handle = async ({ event, resolve }) => {
    const token = event.request.headers.get("Authorization")?.split("Bearer ")[1];
    if (!token) {
        console.info("No token found");
        event.locals.user = null;
        return await resolve(event);
    }

    const app = initializeApp({
        credential: credential.cert(serviceAccount as ServiceAccount),
    });

    app
        .auth()
        .verifyIdToken(token)
        .then((decodedToken) => {
            console.log(decodedToken);
        });

    const response = await resolve(event);
    return response;
};
