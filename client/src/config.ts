import { checkEnv } from "$lib/env.util";

export const Config = {
    VITE_API_URL: checkEnv(import.meta.env.VITE_API_URL),
    VITE_NODE_ENV: checkEnv(import.meta.env.VITE_NODE_ENV),
    VITE_FIREBASE_API_KEY: checkEnv(import.meta.env.VITE_FIREBASE_API_KEY),
    VITE_FIREBASE_AUTH_DOMAIN: checkEnv(import.meta.env.VITE_FIREBASE_AUTH_DOMAIN),
    DAY_FORMAT: 'YYYY-MM-DD',
    MONTH_FORMAT: 'YYYY-MM',
    DAYTIME_FORMAT: 'YYYY-MM-DDTHH:mm:ss',
    DAYTIME_FORMAT_SHORT: 'MM-DD HH:mm',
    STRING_MEDIUM: 200,
    STRING_LONG: 1000,
    MAX_FILE_SIZE: 10000000,
    USER_SESSION_EXPIRE_IN_HOURS: 24,
};
