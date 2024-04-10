<script>
    import { checkElement } from "$lib/helpers";
    import { scale } from "svelte/transition";

    /** @type {string} */
    export let email;
    /** @type {string} */
    export let avatar;

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
        class="-m-1.5 flex items-center p-1.5"
        aria-expanded={open}
        aria-haspopup="true"
        on:click={() => (open = !open)}
    >
        <span class="sr-only">Open user menu</span>
        {#if avatar}
            <img class="h-8 w-8 rounded-full" src={avatar} alt="Avatar" />
        {:else}
            <svg
                class="h-8 w-8 rounded-full"
                fill="currentColor"
                viewBox="0 0 24 24"
            >
                <path
                    d="M24 20.993V24H0v-2.996A14.977 14.977 0 0112.004 15c4.904 0 9.26 2.354 11.996 5.993zM16.002 8.999a4 4 0 11-8 0 4 4 0 018 0z"
                />
            </svg>
        {/if}

        <span class="hidden lg:flex lg:items-center">
            <span
                class="ml-4 text-sm font-semibold leading-6"
                aria-hidden="true"
            >
                Account
            </span>
            <svg
                class="ml-2 h-5 w-5 text-gray-400"
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
        </span>
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
            class="absolute right-0 z-10 mt-2 min-w-[224px] origin-top-right rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none"
            tabindex="-1"
            role="menu"
            aria-orientation="vertical"
            aria-labelledby="menu-button"
        >
            <div class="divide-y divide-gray-100">
                <!-- Active: "bg-gray-100 text-gray-900", Not Active: "text-gray-700" -->

                <div class="py-1" role="none">
                    <span
                        class="block px-4 py-2 text-sm text-gray-700"
                        role="menuitem"
                        tabindex="-1"
                        id="menu-item-0"
                    >
                        Signed in as <strong>{email}</strong>
                    </span>
                </div>
                <div class="py-1" role="none">
                    <a
                        on:mouseover={() => (active = 0)}
                        on:focus={() => (active = 0)}
                        on:click={() => (open = false)}
                        href="https://github.com/mpiorowski/sgsg"
                        target="_blank"
                        class="block px-4 py-2 text-sm text-gray-700
                    {active === 0 ? 'bg-gray-100 text-gray-900' : ''}"
                        role="menuitem"
                        tabindex="-1"
                        id="menu-item-0"
                    >
                        GitHub
                    </a>
                    <a
                        on:mouseover={() => (active = 1)}
                        on:focus={() => (active = 1)}
                        on:click={() => (open = false)}
                        href="https://www.upsend.app/"
                        target="_blank"
                        class="block px-4 py-2 text-sm text-gray-700
                    {active === 1 ? 'bg-gray-100 text-gray-900' : ''}"
                        role="menuitem"
                        tabindex="-1"
                        id="menu-item-1"
                    >
                        UpSend
                    </a>
                </div>
                <div class="py-1" role="none">
                    <button
                        on:click={() => (window.location.href = "/auth")}
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
                </div>
            </div>
        </div>
    {/if}
</div>
