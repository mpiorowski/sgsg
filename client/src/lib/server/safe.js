import { logger } from "$lib/server/logger";

/**
 * @param {Promise<T> | Function} promiseOrFunc
 * @returns {Promise<import("./safe").Safe<T>> | import("./safe").Safe<T>}
 * @template T
 */
export function safe(promiseOrFunc) {
    if (promiseOrFunc instanceof Promise) {
        return safeAsync(promiseOrFunc);
    }
    return safeSync(promiseOrFunc);
}

/**
 * @param {Promise<T>} promise
 * @returns {Promise<import("./safe").Safe<T>>}
 * @template T
 * @private
 */
async function safeAsync(promise) {
    try {
        const data = await promise;
        return { data, success: true };
    } catch (e) {
        logger.error(e);
        if (e instanceof Error) {
            return { success: false, error: e.message };
        }
        return { success: false, error: "Something went wrong" };
    }
}

/**
 * @param {Function} func
 * @returns {import("./safe").Safe<T>}
 * @template T
 * @private
 */
function safeSync(func) {
    try {
        const data = func();
        return { data, success: true };
    } catch (e) {
        logger.error(e);
        if (e instanceof Error) {
            return { success: false, error: e.message };
        }
        return { success: false, error: "Something went wrong" };
    }
}
