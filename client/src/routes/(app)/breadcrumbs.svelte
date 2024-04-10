<script>
    import { page } from "$app/stores";

    $: url = $page.url.pathname.split("/").filter((el) => el !== "");
</script>

<nav class="flex" aria-label="Breadcrumb">
    <ol role="list" class="flex items-center space-x-4">
        <li>
            <div>
                <a href="/" class="text-gray-400 hover:text-gray-400">
                    <svg
                        class="h-5 w-5 flex-shrink-0"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                        aria-hidden="true"
                    >
                        <path
                            fill-rule="evenodd"
                            d="M9.293 2.293a1 1 0 011.414 0l7 7A1 1 0 0117 11h-1v6a1 1 0 01-1 1h-2a1 1 0 01-1-1v-3a1 1 0 00-1-1H9a1 1 0 00-1 1v3a1 1 0 01-1 1H5a1 1 0 01-1-1v-6H3a1 1 0 01-.707-1.707l7-7z"
                            clip-rule="evenodd"
                        />
                    </svg>
                    <span class="sr-only">Home</span>
                </a>
            </div>
        </li>
        {#each url as el, index}
            <li>
                <div class="flex items-center">
                    <svg
                        class="h-5 w-5 flex-shrink-0 text-gray-400"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                        aria-hidden="true"
                    >
                        <path
                            fill-rule="evenodd"
                            d="M7.21 14.77a.75.75 0 01.02-1.06L11.168 10 7.23 6.29a.75.75 0 111.04-1.08l4.5 4.25a.75.75 0 010 1.08l-4.5 4.25a.75.75 0 01-1.06-.02z"
                            clip-rule="evenodd"
                        />
                    </svg>
                    <a
                        href="/{url.slice(0, index + 1).join("/")}"
                        class="ml-4 text-sm font-medium text-gray-400 hover:text-gray-400"
                    >
                        {String(el).charAt(0).toUpperCase() +
                            String(el).slice(1)}
                    </a>
                </div>
            </li>
        {/each}
    </ol>
</nav>
