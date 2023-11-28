<script>
    import { slide } from "svelte/transition";
    import Avatar from "./Avatar.svelte";
    import { page } from "$app/stores";
    import LogoIcon from "$lib/icons/LogoIcon.svelte";

    /** @type {import("./$types").PageData} */
    export let data;

    let mobileMenuOpen = false;
    $: current = $page.url.pathname;
</script>

<nav class="bg-gray-800">
    <div class="mx-auto max-w-7xl sm:px-6 lg:px-8">
        <div class="border-b border-gray-700">
            <div class="flex h-16 items-center justify-between px-4 sm:px-0">
                <div class="flex items-center">
                    <div class="flex-shrink-0">
                        <a
                            href="/"
                            aria-label="Home"
                            class="inline-flex h-8 w-8 items-center justify-center text-indigo-600"
                        >
                            <span class="sr-only">Home</span>
                            <LogoIcon />
                        </a>
                    </div>
                    <div class="hidden md:block">
                        <div class="ml-10 flex items-baseline space-x-4">
                            <!-- Current: "bg-gray-900 text-white", Default: "text-gray-300 hover:bg-gray-700 hover:text-white" -->
                            <a
                                href="/profile"
                                class="rounded-md px-3 py-2 text-sm font-medium
                                focus:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:ring-offset-gray-800
                                {current.startsWith('/profile')
                                    ? 'bg-gray-900 text-white'
                                    : 'text-gray-300 hover:bg-gray-700 hover:text-white'}"
                                aria-current={current.startsWith("/profile")
                                    ? "page"
                                    : undefined}
                            >
                                Profile
                            </a>
                            <a
                                href="/notes"
                                class="rounded-md px-3 py-2 text-sm font-medium
                                focus:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:ring-offset-gray-800
                                {current.startsWith('/notes')
                                    ? 'bg-gray-900 text-white'
                                    : 'text-gray-300 hover:bg-gray-700 hover:text-white'}"
                                aria-current={current.startsWith("/notes")
                                    ? "page"
                                    : undefined}
                            >
                                Notes
                            </a>
                        </div>
                    </div>
                </div>
                <div class="hidden md:block">
                    <div class="ml-4 flex items-center md:ml-6">
                        <!-- Profile dropdown -->
                        <div class="relative ml-3 flex">
                            <Avatar avatar={data.user.avatar} />
                        </div>
                    </div>
                </div>
                <div class="-mr-2 flex md:hidden">
                    <!-- Mobile menu button -->
                    <button
                        on:click={() => (mobileMenuOpen = !mobileMenuOpen)}
                        type="button"
                        class="relative inline-flex items-center justify-center rounded-md bg-gray-800 p-2 text-gray-400 hover:bg-gray-700 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800"
                        aria-controls="mobile-menu"
                        aria-expanded="false"
                    >
                        <span class="absolute -inset-0.5" />
                        <span class="sr-only">Open main menu</span>
                        <!-- Menu open: "hidden", Menu closed: "block" -->
                        <svg
                            class="block h-6 w-6"
                            fill="none"
                            viewBox="0 0 24 24"
                            stroke-width="1.5"
                            stroke="currentColor"
                            aria-hidden="true"
                        >
                            <path
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
                            />
                        </svg>
                        <!-- Menu open: "block", Menu closed: "hidden" -->
                        <svg
                            class="hidden h-6 w-6"
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
            </div>
        </div>
    </div>

    <!-- Mobile menu, show/hide based on menu state. -->
    {#if mobileMenuOpen}
        <div
            transition:slide
            class="border-b border-gray-700 md:hidden"
            id="mobile-menu"
        >
            <div class="space-y-1 px-2 py-3 sm:px-3">
                <!-- Current: "bg-gray-900 text-white", Default: "text-gray-300 hover:bg-gray-700 hover:text-white" -->
                <a
                    href="/profile"
                    class="block rounded-md px-3 py-2 text-base font-medium
                    {current.startsWith('/profile')
                        ? 'bg-gray-900 text-white'
                        : 'text-gray-300 hover:bg-gray-700 hover:text-white'}"
                    aria-current={current.startsWith("/profile")
                        ? "page"
                        : undefined}
                >
                    Profile
                </a>
                <a
                    href="/notes"
                    class="block rounded-md px-3 py-2 text-base font-medium
                    {current.startsWith('/notes')
                        ? 'bg-gray-900 text-white'
                        : 'text-gray-300 hover:bg-gray-700 hover:text-white'}"
                    aria-current={current.startsWith("/notes")
                        ? "page"
                        : undefined}
                >
                    Notes
                </a>
            </div>
            <div class="border-t border-gray-700 pb-3 pt-4">
                <div class="flex items-center px-5">
                    <div class="flex-shrink-0">
                        <img
                            class="h-10 w-10 rounded-full"
                            src={data.user.avatar}
                            alt="Avatar"
                        />
                    </div>
                    <div class="ml-3">
                        <div
                            class="text-base font-medium leading-none text-white"
                        >
                            {data.user.email}
                        </div>
                    </div>
                </div>
                <div class="mt-3 space-y-1 px-2">
                    <a
                        href="https://github.com/mpiorowski/sgsg"
                        target="_blank"
                        class="block rounded-md px-3 py-2 text-base font-medium text-gray-400 hover:bg-gray-700 hover:text-white"
                    >
                        GitHub
                    </a>
                    <a
                        href="https://www.upsend.app/"
                        target="_blank"
                        class="block rounded-md px-3 py-2 text-base font-medium text-gray-400 hover:bg-gray-700 hover:text-white"
                    >
                        UpSend
                    </a>
                    <button
                        class="block rounded-md px-3 py-2 text-base font-medium text-gray-400 hover:bg-gray-700 hover:text-white"
                        on:click={() => {
                            window.location.href = "/";
                        }}
                    >
                        Sign out
                    </button>
                </div>
            </div>
        </div>
    {/if}
</nav>
