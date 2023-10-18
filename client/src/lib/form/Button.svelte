<script>
    /** @type {'button' | 'submit' | 'reset'} */
    export let type = "submit";
    /** @type {'primary' | 'secondary' | 'link' | 'danger'} */
    export let variant = "primary";
    /** @type {string | undefined} */
    export let form = undefined;
    /** @type {string} */
    export let href = "";

    let className = "";
    export { className as class };
</script>

{#if !href}
    <button
        on:click
        {form}
        {type}
        class="group inline-flex w-full items-center justify-center gap-x-1.5 rounded-md px-3 py-2 text-sm font-semibold transition focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2
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
        <slot />
    </button>
{/if}

{#if href}
    <a
        {href}
        class="group inline-flex w-full items-center justify-center gap-x-1.5 rounded-md px-3 py-2 text-sm font-semibold no-underline transition focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2
        {variant === 'primary' &&
            'bg-indigo-500 text-white shadow-sm hover:bg-indigo-400 focus-visible:outline-indigo-600'}
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
