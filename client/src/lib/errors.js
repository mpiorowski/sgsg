/**
 * Extract errors from a fields object
 * @param {{ field: string, tag: string }[] | undefined} fields
 * @param {string} field
 * @returns {string}
 */
export function extractError(fields, field) {
    if (!fields) {
        return "";
    }
    for (const { field: key, tag: value } of fields) {
        if (key.toLowerCase() === field.toLowerCase()) {
            for (const [k, v] of Object.entries(errorMessages)) {
                if (k === value) {
                    return v;
                }
            }
            return "Unknown error";
        }
    }
    return "";
}

const errorMessages = {
    required: "This field is required",
    min: "A minimum of 3 characters is required",
    max: "A maximum of 1000 characters is allowed",
    uuid: "Please enter a valid UUID",
    email: "Please enter a valid email address",
    url: "Please enter a valid URL",
    number: "Please enter a valid number",
    numeric: "Please enter a valid number",
    gt: "Please enter a number greater than 0",
};
