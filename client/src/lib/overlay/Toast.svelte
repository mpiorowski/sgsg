<script>
    import { toastStore } from "$lib/overlay/toast";
    import { cubicOut } from "svelte/easing";
    import { fade, fly } from "svelte/transition";

    /** @type {import("$lib/types").Toast} */
    export let toast;
</script>

<!--
      Notification panel, dynamically insert this into the live region when it needs to be displayed

      Entering: "transform ease-out duration-300 transition"
        From: "translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-2"
        To: "translate-y-0 opacity-100 sm:translate-x-0"
      Leaving: "transition ease-in duration-100"
        From: "opacity-100"
        To: "opacity-0"
    -->
<div
    class="pointer-events-auto w-full max-w-sm overflow-hidden rounded-lg bg-white shadow-lg ring-1 ring-black ring-opacity-5"
    in:fly={{ duration: 300, easing: cubicOut, x: 10 }}
    out:fade={{ duration: 100 }}
    role="alert"
>
    <div class="p-4">
        <div class="flex items-start">
            <div class="flex-shrink-0">
                {#if toast.type === "success"}
                    <svg
                        class="h-6 w-6 text-green-400"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                    >
                        <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14" />
                        <polyline points="22 4 12 14.01 9 11.01" />
                    </svg>
                {:else if toast.type === "error"}
                    <svg
                        class="h-6 w-6 text-red-400"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                    >
                        <circle cx="12" cy="12" r="10" />
                        <line x1="15" y1="9" x2="9" y2="15" />
                        <line x1="9" y1="9" x2="15" y2="15" />
                    </svg>
                {:else if toast.type === "warning"}
                    <svg
                        class="h-6 w-6 text-yellow-400"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                    >
                        <circle cx="12" cy="12" r="10" />
                        <line x1="12" y1="8" x2="12" y2="12" />
                        <line x1="12" y1="16" x2="12.01" y2="16" />
                    </svg>
                {:else if toast.type === "info"}
                    <svg
                        class="h-6 w-6 text-blue-400"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                    >
                        <circle cx="12" cy="12" r="10" />
                        <path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3" />
                        <line x1="12" y1="17" x2="12.01" y2="17" />
                    </svg>
                {/if}
            </div>
            <div class="ml-3 w-0 flex-1 pt-0.5">
                <p class="text-sm font-medium text-gray-900">
                    {toast.title}
                </p>
                {#if toast.description}
                    <p class="mt-1 text-sm text-gray-500">
                        {toast.description}
                    </p>
                {/if}
                {#if toast.action}
                    <button
                        class="mt-4 text-sm text-indigo-600 hover:text-indigo-500"
                        aria-hidden="true"
                        on:click={() => {
                            toast.action?.onClick();
                        }}
                    >
                        {toast.action?.label}
                    </button>
                {/if}
            </div>
            <div class="ml-4 flex flex-shrink-0">
                <button
                    type="button"
                    class="inline-flex rounded-md bg-white text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                    on:click={() => {
                        toastStore.update((toasts) => {
                            return toasts.filter((t) => t.id !== toast.id);
                        });
                    }}
                >
                    <span class="sr-only">Close</span>
                    <svg
                        class="h-5 w-5"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                        aria-hidden="true"
                    >
                        <path
                            d="M6.28 5.22a.75.75 0 00-1.06 1.06L8.94 10l-3.72 3.72a.75.75 0 101.06 1.06L10 11.06l3.72 3.72a.75.75 0 101.06-1.06L11.06 10l3.72-3.72a.75.75 0 00-1.06-1.06L10 8.94 6.28 5.22z"
                        />
                    </svg>
                </button>
            </div>
        </div>
    </div>
</div>
