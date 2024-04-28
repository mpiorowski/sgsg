import { logger } from "./logger";
import { safe } from "./safe";

/**
 * API wrapper
 * @param {string} url
 * @param {{
 *  method?: "GET" | "POST" | "PUT" | "DELETE" | "PATCH"
 *  body?: D
 *  metadata?: string
 *  }} options
 *  @returns {Promise<import("./safe").Safe<T>>}
 *  @template T
 *  @template D = T
 */
export async function api(
    url,
    { method = "GET", body = undefined, metadata } = {},
) {
    logger.debug({ url, method, body, metadata }, "API request");
    const res = await safe(
        fetch(url, {
            method,
            headers: {
                "content-type": "application/json",
                ...(metadata ? { authorization: `Bearer ${metadata}` } : {}),
            },
            body: body ? JSON.stringify(body) : null,
        }),
    );

    if (!res.success) {
        return { success: false, error: res.error };
    }

    // check if error response
    if (!res.data.ok) {
        return { success: false, error: res.data.statusText };
    }

    // check if empty response
    if (res.data.status === 204) {
        const empty = /** @type {T} */ ({});
        return { success: true, data: empty };
    }
    // check if invalid response
    if (!res.data.headers.get("content-type")?.includes("application/json")) {
        return { success: false, error: "Response was not JSON" };
    }
    const data = await res.data.json();
    return { success: true, data };
}
