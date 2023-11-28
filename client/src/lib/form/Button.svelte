<script>
    /** @type {'button' | 'submit' | 'reset'} */
    export let type = "submit";
    /** @type {'primary' | 'secondary' | 'link' | 'danger'} */
    export let variant = "primary";
    /** @type {string | undefined} */
    export let form = undefined;
    /** @type {string} */
    export let href = "";
    /** @type {boolean} */
    export let loading = false;

    let className = "";
    export { className as class };
</script>

{#if !href}
    <button
        on:click
        {form}
        {type}
        disabled={loading}
        class="group inline-flex items-center justify-center gap-x-1.5 rounded-md px-3 py-2 text-sm font-semibold transition focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2
        {variant === 'primary' &&
            'bg-indigo-600 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline-indigo-600'}
        {variant === 'secondary' &&
            'bg-white text-gray-900 shadow-sm ring-1 ring-gray-300 hover:bg-gray-50 focus-visible:outline-gray-300'}
        {variant === 'link' &&
            'text-indigo-600 hover:text-indigo-500 focus-visible:outline-indigo-600'}
        {variant === 'danger' &&
            'bg-red-600 text-white shadow-sm hover:bg-red-500 focus-visible:outline-red-600'}
        {className}
        "
    >
        {#if loading}
            <svg
                class="h-5 w-5 animate-spin text-white"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
            >
                <circle
                    class="opacity-25"
                    cx="12"
                    cy="12"
                    r="10"
                    stroke="currentColor"
                    stroke-width="4"
                />
                <path
                    class="opacity-75"
                    fill="currentColor"
                    d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                />
            </svg>
        {:else}
            <slot name="icon" />
        {/if}
        <slot />
    </button>
{/if}

{#if href}
    <a
        {href}
        class="group inline-flex items-center justify-center gap-x-1.5 rounded-md px-3 py-2 text-sm font-semibold no-underline transition focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2
        {variant === 'primary' &&
            'bg-indigo-600 text-white shadow-sm hover:bg-indigo-400 focus-visible:outline-indigo-600'}
        {variant === 'secondary' &&
            'bg-white text-gray-900 shadow-sm ring-1 ring-gray-300 hover:bg-gray-50 focus-visible:outline-gray-300'}
        {variant === 'link' &&
            'text-indigo-600 hover:text-indigo-500 focus-visible:outline-indigo-600'}
        {variant === 'danger' &&
            'bg-red-600 text-white shadow-sm hover:bg-red-500 focus-visible:outline-red-600'}
        {className}
        "
    >
        <slot />
    </a>
{/if}
