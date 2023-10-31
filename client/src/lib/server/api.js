import { SERVER_HTTP, UPSEND_KEY } from "$env/static/private";
import { logger } from "./logger";

/**
 * Fetch data from API
 * @param {{
 *  method?: "GET" | "POST" | "PUT" | "DELETE",
 *  url: string,
 *  body?: object,
 *  token: string
 * }} options
 * @returns {Promise<import("$lib/safe").Safe<T>>}
 * @template T
 */
export async function api({ method = "GET", url, body, token }) {
    try {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");
        headers.append("Authorization", `Bearer ${token}`);

        const response = await fetch(`${SERVER_HTTP}${url}`, {
            method,
            headers,
            body: JSON.stringify(body),
        });
        if (!response.ok) {
            throw new Error(await response.text());
        }
        const data = await response.json();

        return { error: false, data };
    } catch (error) {
        logger.error(error);
        return { error: true, msg: "Error during fetch" };
    }
}

/**
 * Upsend API
 * @param {{
 *  method?: "GET" | "POST" | "DELETE",
 *  url: string,
 *  file?: File,
 * }} options
 * @returns {Promise<import("$lib/safe").Safe<T>>}
 * @template T
 */
export async function upsendApi({ method = "GET", url, file }) {
    try {
        const headers = new Headers();
        headers.append("Authorization", `Bearer ${UPSEND_KEY}`);

        let formData = null;
        if (file) {
            formData = new FormData();
            formData.append("file", file);
        }
        const response = await fetch("https://api.upsend.app" + url, {
            method,
            headers,
            body: formData,
        });
        if (!response.ok) {
            throw new Error(await response.text());
        }
        if (response.status === 204) {
            const empty = /** @type {T} */ ("");
            return { error: false, data: empty };
        }
        const data = await response.json();

        return { error: false, data };
    } catch (error) {
        logger.error(error);
        return {
            error: true,
            msg: "Error using Upsend API",
        };
    }
}
