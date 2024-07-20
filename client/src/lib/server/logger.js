import { env } from "$env/dynamic/private";
import pino from "pino";

function get_pino_config() {
    if (env.TARGET === "development") {
        return {
            transport: {
                target: "pino-pretty",
                options: {
                    colorize: true,
                },
            },
            level: "debug",
        };
    } else {
        return {
            level: "info",
        };
    }
}

export const logger = pino(get_pino_config());

/**
 * Measure the performance
 * @param {string} name - The name of the performance measurement
 * @returns {() => void} - The end function
 */
export function perf(name) {
    if (env.TARGET === "production") {
        return () => {
            // do nothing
        };
    }
    const start = performance.now();

    /**
     * End the performance measurement
     * @returns {void}
     */
    function end() {
        const duration = performance.now() - start;
        logger.info(`${name}: ${duration.toFixed(4)}ms`);
    }

    return end;
}
