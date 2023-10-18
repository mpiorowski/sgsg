<script>
    import { browser } from "$app/environment";
    import Button from "$lib/form/Button.svelte";
    import { checkElement, generateId } from "$lib/helpers";
    import XIcon from "$lib/icons/XIcon.svelte";
    import { fade, fly } from "svelte/transition";

    /** @type {boolean} */
    export let open = false;
    /** @type {string} */
    export let title;
    /** @type {"right" | "left"} */
    export let position = "right";

    /** @type {HTMLElement | undefined} */
    let previous = undefined;

    $: if (!open) {
        previous?.focus({ preventScroll: true });
    }
    $: if (browser) {
        document.body.classList.toggle("no-scroll", open);
        const scroll = document.getElementById("scroll");
        scroll?.classList.toggle("scroll", open);
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
    class="relative z-10"
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
        class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"
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
                    class="pointer-events-auto w-screen max-w-xl"
                >
                    <div
                        class="flex h-full flex-col divide-y divide-gray-200 bg-white shadow-xl"
                    >
                        <div
                            class="flex min-h-0 flex-1 flex-col overflow-y-scroll py-6"
                        >
                            <div class="px-4 sm:px-6">
                                <div class="flex items-start justify-between">
                                    <h2
                                        class="text-base font-semibold leading-6 text-gray-900"
                                        id="drawer-title-{id}"
                                    >
                                        {title}
                                    </h2>
                                    <div class="ml-3 flex h-7 items-center">
                                        <button
                                            on:click={() => (open = false)}
                                            type="button"
                                            class="relative rounded-md bg-white text-gray-400 hover:text-gray-500 focus-visible:ring-2 focus-visible:ring-indigo-500"
                                        >
                                            <span class="absolute -inset-2.5" />
                                            <span class="sr-only">
                                                Close panel
                                            </span>
                                            <XIcon />
                                        </button>
                                    </div>
                                </div>
                            </div>
                            <div class="relative mt-6 flex-1 px-4 sm:px-6">
                                <slot />
                            </div>
                        </div>
                        <div class="flex justify-end px-4 py-4">
                            <div class="inline-flex gap-4">
                                <Button
                                    on:click={() => (open = false)}
                                    type="button"
                                    variant="secondary"
                                >
                                    Cancel
                                </Button>
                                <slot name="submit" />
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
