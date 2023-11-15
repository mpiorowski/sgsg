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
        return { data, error: false };
    } catch (e) {
        logger.error(e);
        if (e instanceof Error) {
            return { error: true, msg: e.message };
        }
        return { error: true, msg: "Something went wrong" };
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
        return { data, error: false };
    } catch (e) {
        logger.error(e);
        if (e instanceof Error) {
            return { error: true, msg: e.message };
        }
        return { error: true, msg: "Something went wrong" };
    }
}

/**
 * Callback function for handling gRPC responses safely.
 *
 * @template T - The type of data expected in the response.
 *
 * @param {(value: import("./safe").Safe<T>) => void} res - The callback function to handle the response.
 * @returns {(err: import("@grpc/grpc-js").ServiceError | null, data: T | undefined) => void} - A callback function to be used with gRPC response handling.
 */
export function grpcSafe(res) {
    /**
     * Handles the gRPC response and calls the provided callback function safely.
     *
     * @param {import("@grpc/grpc-js").ServiceError | null} err - The error, if any, returned in the response.
     * @param {T | undefined} data - The data returned in the response.
     */
    return (err, data) => {
        if (err) {
            logger.error(err);
            if (err.code === 3) {
                let fields = [];
                try {
                    fields = JSON.parse(err.details);
                } catch (e) {
                    return res({
                        error: true,
                        msg: err?.message || "Something went wrong",
                    });
                }

                return res({
                    error: true,
                    msg: "Invalid argument",
                    fields: fields,
                });
            }
            return res({
                error: true,
                msg: err?.message || "Something went wrong",
            });
        }
        if (!data) {
            logger.error("No data returned");
            return res({
                error: true,
                msg: "No data returned",
            });
        }
        res({ data, error: false });
    };
}
