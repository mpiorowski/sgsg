import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";

const firebaseConfig = {
    apiKey: "AIzaSyADPQAnPOZFTexgfZ0kE3cY0PfaxWvv3Fc",
    authDomain: "go-svelte-grpc.firebaseapp.com",
}

const app = initializeApp(firebaseConfig);
export const clientAuth = getAuth(app);
