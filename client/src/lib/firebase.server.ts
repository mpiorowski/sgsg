import { credential, initializeApp, type ServiceAccount } from "firebase-admin";
import serviceAccount from "../../../serviceAccount.json";

const server = initializeApp({
    credential: credential.cert(serviceAccount as ServiceAccount),
}, "server");
export const serverAuth = server.auth();
