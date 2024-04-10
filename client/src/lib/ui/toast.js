import { writable } from "svelte/store";
import { generateId } from "$lib/helpers";

/**
 * @typedef {{
 * id: string;
 * type: "success" | "error" | "warning" | "info";
 * title: string;
 * description: string;
 * duration: number;
 * action?: {label: string, onClick: () => void};
 * }} Toast
 */
export {};

/** @type {import("svelte/store").Writable<Toast[]>} */
export const toastStore = writable([]);

/**
 * Show a toast
 * @param {Toast} toast
 * @returns {void}
 */
export function showToast(toast) {
    toastStore.update((toasts) => [...toasts, toast]);
    setTimeout(() => {
        toastStore.update((toasts) => {
            return toasts.filter((t) => t.id !== toast.id);
        });
    }, toast.duration);
}

/**
 * Toast helper
 * @type {{
 * success: (title: string, description?: string) => void,
 * error: (title: string, description?: string) => void,
 * warning: (title: string, description?: string) => void,
 * info: (title: string, description?: string) => void
 * }}
 */
export const toast = {
    success: (title, description = "") =>
        showToast({
            id: generateId(),
            title,
            description,
            type: "success",
            duration: 5000,
        }),
    error: (title, description = "") =>
        showToast({
            id: generateId(),
            title,
            description,
            type: "error",
            duration: 5000,
        }),
    warning: (title, description = "") =>
        showToast({
            id: generateId(),
            title,
            description,
            type: "warning",
            duration: 5000,
        }),
    info: (title, description = "") =>
        showToast({
            id: generateId(),
            title,
            description,
            type: "info",
            duration: 5000,
        }),
};
