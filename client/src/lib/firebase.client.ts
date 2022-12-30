import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";
import { Config } from "src/config";

const firebaseConfig = {
    apiKey: Config.VITE_FIREBASE_API_KEY,
    authDomain: Config.VITE_FIREBASE_AUTH_DOMAIN,
};

const app = initializeApp(firebaseConfig);
export const clientAuth = getAuth(app);
