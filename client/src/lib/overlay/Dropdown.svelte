<script>
    import { checkElement } from "$lib/helpers";
    import { scale } from "svelte/transition";
    /** @type {boolean} */
    export let rounded = false;

    /** @type {boolean} */
    let open = false;
    /** @type {number} */
    let active = 0;

    /**
     * @param {HTMLElement} node
     * @returns {{ destroy(): void }}
     */
    function portal(node) {
        const previous = checkElement(document.activeElement);

        /** @type {NodeListOf<HTMLElement>} */
        const focusableElements = node.querySelectorAll(
            'a[href], button, textarea, input[type="text"], input[type="radio"], input[type="checkbox"], select',
        );
        const first = focusableElements[0];
        const items = focusableElements.length;
        first?.focus({ preventScroll: true });

        /**
         * @param {KeyboardEvent} event
         * @returns {void}
         */
        function handleKeydown(event) {
            if (event.key === "ArrowDown") {
                event.preventDefault();
                active = (active + 1) % items;
                focusableElements[active]?.focus({ preventScroll: true });
            }
            if (event.key === "ArrowUp") {
                event.preventDefault();
                active = (active - 1 + items) % items;
                focusableElements[active]?.focus({ preventScroll: true });
            }
            if (event.key === "Escape") {
                open = false;
                previous?.focus({ preventScroll: true });
            }
            if (event.key === "Tab") {
                event.preventDefault();
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
            if (
                !node.contains(event.target) &&
                !previous?.contains(event.target)
            ) {
                open = false;
            }
        }

        node.addEventListener("keydown", handleKeydown);
        document.addEventListener("click", handleClickOutside);

        return {
            destroy() {
                node.removeEventListener("keydown", handleKeydown);
                document.removeEventListener("click", handleClickOutside);
            },
        };
    }
</script>

<div class="relative inline-block text-left">
    <button
        type="button"
        id="menu-button"
        class="flex justify-center items-center focus-visible:ring-2 focus-visible:ring-indigo-500 focus-visible:ring-offset-2
        {rounded ? 'rounded-full' : ''}"
        aria-expanded={open}
        aria-haspopup="true"
        on:click={() => (open = !open)}
    >
        <!-- class="inline-flex w-full justify-center gap-x-1.5 rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50" -->
        <slot name="button" />
        <!--
            Options
            <svg
                class="-mr-1 h-5 w-5 text-gray-400"
                viewBox="0 0 20 20"
                fill="currentColor"
                aria-hidden="true"
            >
                <path
                    fill-rule="evenodd"
                    d="M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z"
                    clip-rule="evenodd"
                />
            </svg>
        -->
    </button>

    {#if open}
        <!--
    Dropdown menu, show/hide based on menu state.

    Entering: "transition ease-out duration-100"
      From: "transform opacity-0 scale-95"
      To: "transform opacity-100 scale-100"
    Leaving: "transition ease-in duration-75"
      From: "transform opacity-100 scale-100"
      To: "transform opacity-0 scale-95"
  -->
        <div
            use:portal
            in:scale={{ duration: 100, start: 0.95, opacity: 0 }}
            out:scale={{ duration: 75, start: 0.95, opacity: 0 }}
            class="absolute right-0 z-10 mt-2 w-56 origin-top-right rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none"
            tabindex="-1"
            role="menu"
            aria-orientation="vertical"
            aria-labelledby="menu-button"
        >
            <div class="py-1" role="none">
                <!-- Active: "bg-gray-100 text-gray-900", Not Active: "text-gray-700" -->
                <a
                    on:mouseover={() => (active = 0)}
                    on:focus={() => (active = 0)}
                    href="/account"
                    class="block px-4 py-2 text-sm text-gray-700
                    {active === 0 ? 'bg-gray-100 text-gray-900' : ''}"
                    role="menuitem"
                    tabindex="-1"
                    id="menu-item-0"
                >
                    Account settings
                </a>
                <a
                    on:mouseover={() => (active = 1)}
                    on:focus={() => (active = 1)}
                    href="/support"
                    class="block px-4 py-2 text-sm text-gray-700
                    {active === 1 ? 'bg-gray-100 text-gray-900' : ''}"
                    role="menuitem"
                    tabindex="-1"
                    id="menu-item-1"
                >
                    Support
                </a>
                <a
                    on:mouseover={() => (active = 2)}
                    on:focus={() => (active = 2)}
                    href="/license"
                    class="block px-4 py-2 text-sm text-gray-700
                    {active === 2 ? 'bg-gray-100 text-gray-900' : ''}"
                    role="menuitem"
                    tabindex="-1"
                    id="menu-item-2"
                >
                    License
                </a>
                <form method="POST" action="#" role="none">
                    <button
                        on:mouseover={() => (active = 3)}
                        on:focus={() => (active = 3)}
                        type="submit"
                        class="block w-full px-4 py-2 text-left text-sm text-gray-700
                        {active === 3 ? 'bg-gray-100 text-gray-900' : ''}"
                        role="menuitem"
                        tabindex="-1"
                        id="menu-item-3"
                    >
                        Sign out
                    </button>
                </form>
            </div>
        </div>
    {/if}
</div>
