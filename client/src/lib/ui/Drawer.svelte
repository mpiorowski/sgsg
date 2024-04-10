<script>
    import { browser } from "$app/environment";
    import { checkElement, generateId } from "$lib/helpers";
    import { fade, fly } from "svelte/transition";

    /** @type {boolean} */
    export let open;
    /** @type {string} */
    export let title = "";
    /** @type {"right" | "left"} */
    export let position = "right";
    /** @type {() => void} */
    export let close;

    /** @type {HTMLElement | undefined} */
    let previous = undefined;

    $: if (!open) {
        previous?.focus({ preventScroll: true });
    }

    $: if (browser) {
        if (open) {
            document.body.classList.add("no-scroll");
        } else {
            document.body.classList.remove("no-scroll");
        }
    }

    /**
     * @see Every change here should must be added to Modal.svelte
     * @link ./Modal.svelte
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
                close();
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
                close();
            }
        }
        document.addEventListener("keydown", handleKeydown);
        document.addEventListener("mousedown", handleClickOutside);

        return {
            destroy() {
                document.removeEventListener("keydown", handleKeydown);
                document.removeEventListener("mousedown", handleClickOutside);
            },
        };
    }

    const id = generateId();
</script>

<div
    class="relative z-40"
    aria-labelledby="drawer-title-{id}"
    role="dialog"
    aria-modal="true"
>
    <!--
    Background backdrop, show/hide based on slide-over state.

    Entering: "ease-in-out duration-500"
      From: "opacity-0"
      To: "opacity-100"
    Leaving: "ease-in-out duration-500"
      From: "opacity-100"
      To: "opacity-0"
  -->
    <div
        transition:fade
        class="fixed inset-0 bg-gray-900 bg-opacity-75 transition-opacity"
    />

    <div class="fixed inset-0 overflow-hidden">
        <div class="absolute inset-0 overflow-hidden">
            <div
                class="pointer-events-none fixed inset-y-0 flex max-w-full
                {position === 'right' ? 'right-0 pl-10' : 'left-0 pr-10'}"
            >
                <!--
          Slide-over panel, show/hide based on slide-over state.

          Entering: "transform transition ease-in-out duration-500 sm:duration-700"
            From: "translate-x-full"
            To: "translate-x-0"
          Leaving: "transform transition ease-in-out duration-500 sm:duration-700"
            From: "translate-x-0"
            To: "translate-x-full"
        -->
                <div
                    use:portal
                    in:fly={{
                        x: position === "right" ? "100%" : "-100%",
                        duration: 400,
                        opacity: 100,
                    }}
                    out:fly={{
                        x: position === "right" ? "100%" : "-100%",
                        duration: 400,
                        opacity: 100,
                    }}
                    class="pointer-events-auto relative w-screen max-w-xl"
                >
                    <!--
                        Close button, show/hide based on slide-over state.

                        Entering: "ease-in-out duration-500"
                          From: "opacity-0"
                          To: "opacity-100"
                        Leaving: "ease-in-out duration-500"
                          From: "opacity-100"
                          To: "opacity-0"
                      -->
                    <div
                        class="absolute top-0 flex pt-4
                            {position === 'left'
                            ? 'right-0 -mr-8 pl-2 sm:-mr-10 sm:pl-4'
                            : 'left-0 -ml-8 pr-2 sm:-ml-10 sm:pr-4'}"
                        transition:fade
                    >
                        <button
                            type="button"
                            class="relative rounded-md text-gray-300 hover:text-white focus:outline-none focus:ring-2 focus:ring-white"
                            on:click={close}
                        >
                            <span class="absolute -inset-2.5"></span>
                            <span class="sr-only">Close panel</span>
                            <svg
                                class="h-6 w-6"
                                fill="none"
                                viewBox="0 0 24 24"
                                stroke-width="1.5"
                                stroke="currentColor"
                                aria-hidden="true"
                            >
                                <path
                                    stroke-linecap="round"
                                    stroke-linejoin="round"
                                    d="M6 18L18 6M6 6l12 12"
                                />
                            </svg>
                        </button>
                    </div>
                    {#if title}
                        <div
                            class="flex h-full flex-col overflow-y-scroll bg-gray-900 py-6 shadow-xl ring-1 ring-white/10"
                        >
                            <div class="px-4 sm:px-6">
                                <h2
                                    class="text-base font-semibold leading-6 text-gray-50"
                                    id="slide-over-title"
                                >
                                    {title}
                                </h2>
                            </div>
                            <div class="relative mt-6 flex-1 px-4 sm:px-6">
                                <slot />
                            </div>
                        </div>
                    {:else}
                        <div
                            class="flex h-full flex-col overflow-y-scroll bg-gray-900 shadow-xl ring-1 ring-white/10"
                        >
                            <slot />
                        </div>
                    {/if}

                    <!--
                    <div
                        class="flex h-full flex-col divide-y overflow-y-scroll divide-gray-200 bg-gray-900 shadow-xl"
                    >
                        <div class="flex min-h-0 flex-1 flex-col py-6">
                            <div class="px-4 sm:px-6">
                                <div class="flex items-start justify-between">
                                    <h2
                                        class="text-base font-semibold leading-6 text-gray-900"
                                        id="drawer-title-{id}"
                                    >
                                        {title}
                                    </h2>
                                </div>
                            </div>
                            <div class="relative mt-6 flex-1 px-4 sm:px-6">
                                <slot />
                            </div>
                        </div>
                        <div class="flex justify-end px-4 py-4">
                            <div class="inline-flex gap-4">
                                <Button
                                    on:click={close}
                                    type="button"
                                    variant="secondary"
                                >
                                    Cancel
                                </Button>
                                <slot name="submit" />
                            </div>
                        </div>
                    </div>
                    -->
                </div>
            </div>
        </div>
    </div>
</div>
