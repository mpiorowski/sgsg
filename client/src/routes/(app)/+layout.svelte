<script>
    import { page } from "$app/stores";
    import Drawer from "$lib/ui/Drawer.svelte";
    import Avatar from "./avatar.svelte";
    import Nav from "./nav.svelte";

    /** @type {import("./$types").LayoutData} */
    export let data;

    /** @type {boolean} */
    let open = false;

    $: current = $page.url.pathname
        .split("/")[1]
        ?.replace(/^\w/, (c) => c.toUpperCase());
</script>

{#if open}
    <Drawer {open} close={() => (open = false)} position="left">
        <Nav close={() => (open = false)} />
    </Drawer>
{/if}

<div class="min-h-full bg-gray-900 font-poppins text-gray-50 antialiased">
    <!-- Static sidebar for desktop -->
    <div
        class="hidden lg:fixed lg:inset-y-0 lg:z-40 lg:flex lg:w-20 lg:flex-col"
    >
        <!-- Sidebar component, swap this element with another sidebar if you like -->
        <Nav />
    </div>

    <!-- Your content -->
    <main class="h-full overflow-auto lg:pl-20">
        <div
            class="sticky top-0 z-30 flex h-16 shrink-0 items-center gap-x-4 border-b border-white/5 bg-gray-900 px-4 shadow-sm lg:px-8"
        >
            <button
                type="button"
                class="-m-2.5 p-2.5 lg:hidden"
                on:click={() => (open = true)}
            >
                <span class="sr-only">Open sidebar</span>
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
                        d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
                    />
                </svg>
            </button>

            <!-- Separator -->
            <div class="h-6 w-px bg-white/10 lg:hidden" aria-hidden="true" />

            <div class="flex flex-1 gap-x-4 self-stretch lg:gap-x-6">
                <h1
                    class="flex flex-1 items-center text-2xl font-bold text-gray-200"
                >
                    {current}
                </h1>
                <div class="flex items-center gap-x-4 lg:gap-x-6">
                    <button
                        type="button"
                        class="-m-2.5 p-2.5 text-gray-400 hover:text-gray-400"
                    >
                        <span class="sr-only">View notifications</span>
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
                                d="M14.857 17.082a23.848 23.848 0 005.454-1.31A8.967 8.967 0 0118 9.75v-.7V9A6 6 0 006 9v.75a8.967 8.967 0 01-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 01-5.714 0m5.714 0a3 3 0 11-5.714 0"
                            />
                        </svg>
                    </button>

                    <!-- Separator -->
                    <div
                        class="hidden lg:block lg:h-6 lg:w-px lg:bg-white/10"
                        aria-hidden="true"
                    />

                    <!-- Profile dropdown -->
                    <Avatar email={data.email} avatar={data.avatar} />
                </div>
            </div>
        </div>

        <div class="container max-w-7xl p-6 lg:py-10 lg:px-16">
            <slot />
        </div>
    </main>
</div>
