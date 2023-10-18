import { SERVER_HTTP } from "$env/static/private";
import { logger } from "./logger";

/**
 * Fetch data from API
 * @param {{
 *  method?: "GET" | "POST" | "PUT" | "DELETE",
 *  url: string,
 *  body?: object,
 *  token?: string
 *  apikey?: string
 * }} options
 * @returns {Promise<import("$lib/safe").Safe<T> & { response: Response | undefined }>}
 * @template T
 */
export async function api({ method = "GET", url, body, token, apikey }) {
    try {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");
        if (token) {
            headers.append("Authorization", `Bearer ${token}`);
        } else if (apikey) {
            headers.append("X-API-KEY", apikey);
        }

        const response = await fetch(`${SERVER_HTTP}${url}`, {
            method,
            headers,
            body: JSON.stringify(body),
        });
        if (!response.ok) {
            logger.error(response.statusText);
            return { error: true, msg: response.statusText, response };
        }
        const data = await response.json();

        return { error: false, data, response };
    } catch (error) {
        logger.error(error);
        return { error: true, msg: "Error during fetch", response: undefined };
    }
}
