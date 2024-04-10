/**
 * @param {Promise<T> | (() => T)} promiseOrFunc
 * @returns {Promise<import("./safe").Safe<T>> | import("./safe").Safe<T>}
 * @template T
 * @public
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
        if (e instanceof Error) {
            return { success: false, error: e.message };
        }
        return { success: false, error: "Something went wrong" };
    }
}

/**
 * @param {() => T} func
 * @returns {import("./safe").Safe<T>}
 * @template T
 * @private
 */
function safeSync(func) {
    try {
        const data = func();
        return { data, success: true };
    } catch (e) {
        if (e instanceof Error) {
            return { success: false, error: e.message };
        }
        return { success: false, error: "Something went wrong" };
    }
}
