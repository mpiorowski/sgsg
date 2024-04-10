<script>
    import { page } from "$app/stores";
    import { pagination } from "./pagination";

    /** @type {number} */
    export let total;
    /** @type {number} */
    export let limit;
    /** @type {number} */
    $: p = Number($page.url.searchParams.get("page")) || 1;

    $: pag = pagination(total, p, limit);
</script>

<div class="flex flex-1 justify-between sm:hidden">
    <a
        href={`?p=${pag.prev}`}
        class="relative inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-400 hover:bg-gray-50"
    >
        Previous
    </a>
    <a
        href={`?p=${pag.next}`}
        class="relative ml-3 inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-400 hover:bg-gray-50"
    >
        Next
    </a>
</div>
<div class="hidden sm:flex sm:flex-1 sm:items-center sm:justify-between">
    <div>
        <p class="text-sm text-gray-400">
            Showing
            <span class="font-medium">{pag.start}</span>
            to
            <span class="font-medium">{pag.end}</span>
            of
            <span class="font-medium">{pag.total}</span>
            results
        </p>
    </div>
    <div>
        <nav
            class="isolate inline-flex -space-x-px rounded-md shadow-sm"
            aria-label="Pagination"
        >
            <a
                href={`?page=${pag.prev}`}
                class="relative inline-flex items-center rounded-l-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-600
                hover:bg-gray-800 focus:z-20 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-0 focus-visible:outline-indigo-600"
            >
                <span class="sr-only">Previous</span>
                <svg
                    class="h-5 w-5"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                    aria-hidden="true"
                >
                    <path
                        fill-rule="evenodd"
                        d="M12.79 5.23a.75.75 0 01-.02 1.06L8.832 10l3.938 3.71a.75.75 0 11-1.04 1.08l-4.5-4.25a.75.75 0 010-1.08l4.5-4.25a.75.75 0 011.06.02z"
                        clip-rule="evenodd"
                    />
                </svg>
            </a>
            <!-- Current: "z-10 bg-indigo-600 text-white focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600", Default: "text-gray-900 ring-1 ring-inset ring-gray-600 hover:bg-gray-50 focus:outline-offset-0" -->
            {#each pag.schema as i}
                {#if i === 0}
                    <span
                        class="relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-400 ring-1 ring-inset ring-gray-600 focus:outline-offset-0"
                    >
                        ...
                    </span>
                {:else}
                    <a
                        href="?page={i}"
                        aria-current={p === i ? "page" : undefined}
                        class="relative z-10 inline-flex items-center px-4 py-2 text-sm font-semibold
                        focus:z-20 focus-visible:outline focus-visible:outline-2 focus-visible:outline-indigo-600
                    {p === i
                            ? 'bg-indigo-600 text-white  focus-visible:outline-offset-2 '
                            : 'text-gray-100 ring-1 ring-inset ring-gray-600 hover:bg-gray-800 focus-visible:outline-offset-0'}"
                    >
                        {i}
                    </a>
                {/if}
            {/each}
            <a
                href={`?page=${pag.next}`}
                class="relative inline-flex items-center rounded-r-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-600
                hover:bg-gray-800 focus:z-20 focus:outline-offset-0 focus-visible:outline focus-visible:outline-2 focus-visible:outline-indigo-600"
            >
                <span class="sr-only">Next</span>
                <svg
                    class="h-5 w-5"
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
            </a>
        </nav>
    </div>
</div>
