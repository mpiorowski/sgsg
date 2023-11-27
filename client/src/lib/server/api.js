import { UPSEND_KEY } from "$env/static/private";
import { logger } from "./logger";

/**
 * Upsend API
 * @param {{
 *  method?: "GET" | "POST" | "DELETE",
 *  url: string,
 *  file?: File,
 *  email?: import("$lib/types").UpsendEmail,
 * }} options
 * @returns {Promise<import("$lib/safe").Safe<T>>}
 * @template T
 */
export async function upsendApi({ method = "GET", url, file, email }) {
    try {
        const headers = new Headers();
        headers.append("Authorization", `Bearer ${UPSEND_KEY}`);

        let body = null;
        if (file) {
            body = new FormData();
            body.append("file", file);
        } else if (email) {
            body = JSON.stringify(email);
            headers.append("Content-Type", "application/json");
        }
        const response = await fetch("https://api.upsend.app" + url, {
            method,
            headers,
            body,
        });
        if (!response.ok) {
            throw new Error(await response.text());
        }
        if (response.status === 204 || response.status === 201) {
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
