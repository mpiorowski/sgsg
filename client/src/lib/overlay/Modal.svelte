<script>
    import { cubicIn, cubicOut } from "svelte/easing";
    import { fade, scale } from "svelte/transition";
    import { checkElement, generateId } from "$lib/helpers";
    import Button from "$lib/form/Button.svelte";

    /** @type {boolean} */
    export let open = false;
    /** @type {string} */
    export let title;
    /** @type {string} */
    export let description;
    /** @type {boolean} */
    export let alert = true;

    /** @type {HTMLElement | undefined} */
    let previous = undefined;

    $: if (!open) {
        previous?.focus({ preventScroll: true });
    }

    /**
     * @see Every change here should must be added to Drawer.svelte
     * @link ./Drawer.svelte
     * @param {HTMLElement} node
     * @returns {{ destroy(): void }}
     */
    function portal(node) {
        previous = checkElement(document.activeElement);

        /** @type {NodeListOf<HTMLElement>} */
        const focusableElements = node.querySelectorAll(
            'a[href], button, textarea, input[type="text"], input[type="radio"], input[type="checkbox"], select',
        );
        const first = focusableElements[0];
        const last = focusableElements[focusableElements.length - 1];
        first?.focus({ preventScroll: true });

        /**
         * @param {KeyboardEvent} event
         * @returns {void}
         */
        function handleKeydown(event) {
            if (event.key === "Escape") {
                open = false;
            }
            if (event.key === "Tab") {
                if (event.shiftKey) {
                    if (document.activeElement === first) {
                        event.preventDefault();
                        last?.focus({ preventScroll: true });
                    }
                } else {
                    if (document.activeElement === last) {
                        event.preventDefault();
                        first?.focus({ preventScroll: true });
                    }
                }
            }
        }

        /**
         * @param {MouseEvent} event
         * @returns {void}
         */
        function handleClickOutside(event) {
            if (!(event.target instanceof Node)) {
                return;
            }
            if (previous?.contains(event.target)) {
                return;
            }
            if (!node.contains(event.target)) {
                open = false;
            }
        }
        document.addEventListener("keydown", handleKeydown);
        document.addEventListener("click", handleClickOutside);

        return {
            destroy() {
                document.removeEventListener("keydown", handleKeydown);
                document.removeEventListener("click", handleClickOutside);
            },
        };
    }

    const id = generateId();
</script>

<div
    class="relative z-10"
    role={alert ? "alertdialog" : "dialog"}
    aria-labelledby="modal-title-{id}"
    aria-describedby="modal-description-{id}"
    aria-modal="true"
>
    <!--
    Background backdrop, show/hide based on modal state.

    Entering: "ease-out duration-300"
      From: "opacity-0"
      To: "opacity-100"
    Leaving: "ease-in duration-200"
      From: "opacity-100"
      To: "opacity-0"
  -->
    <div
        in:fade={{ duration: 300, easing: cubicOut }}
        out:fade={{ duration: 200, easing: cubicIn }}
        class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"
    />

    <div class="fixed inset-0 z-10 w-screen overflow-y-auto">
        <div
            class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0"
        >
            <!--
        Modal panel, show/hide based on modal state.

        Entering: "ease-out duration-300"
          From: "opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
          To: "opacity-100 translate-y-0 sm:scale-100"
        Leaving: "ease-in duration-200"
          From: "opacity-100 translate-y-0 sm:scale-100"
          To: "opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
      -->
            <div
                use:portal
                in:scale={{ duration: 300, easing: cubicOut, start: 0.95 }}
                out:scale={{ duration: 200, easing: cubicIn, start: 0.95 }}
                class="relative transform overflow-hidden rounded-lg bg-white px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg sm:p-6"
            >
                <div class="sm:flex sm:items-start">
                    <div
                        class="mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-red-100 sm:mx-0 sm:h-10 sm:w-10"
                    >
                        <svg
                            class="h-6 w-6 text-red-600"
                            fill="none"
                            viewBox="0 0 24 24"
                            stroke-width="1.5"
                            stroke="currentColor"
                            aria-hidden="true"
                        >
                            <path
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126zM12 15.75h.007v.008H12v-.008z"
                            />
                        </svg>
                    </div>
                    <div class="mt-3 text-center sm:ml-4 sm:mt-0 sm:text-left">
                        <h3
                            class="text-base font-semibold leading-6 text-gray-900"
                            id="modal-title-{id}"
                        >
                            {title}
                        </h3>
                        <div class="mt-2" id="modal-description-{id}">
                            <p class="text-sm text-gray-500">
                                {description}
                            </p>
                        </div>
                    </div>
                </div>
                <div
                    class="mt-5 flex flex-col gap-2 sm:float-right sm:mt-4 sm:inline-flex sm:flex-row-reverse"
                >
                    <Button
                        on:click={() => (open = false)}
                        type="button"
                        variant="secondary"
                    >
                        Cancel
                    </Button>
                    <slot />
                </div>
            </div>
        </div>
    </div>
</div>
