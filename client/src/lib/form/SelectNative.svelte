<script>
    /** @type {string} */
    export let name;
    /** @type {string} */
    export let label;
    /** @type {string | number} */
    export let value;
    /** @type {string} */
    export let error = "";
    /** @type {string} */
    export let helper = "";
</script>

<div class="mb-2">
    <label for={name} class="block text-sm font-medium leading-6">
        {label}
    </label>
    <select
        bind:value
        id={name}
        {name}
        class="mt-2 block w-full rounded-md border-0 bg-gray-800 py-1.5 pl-3 pr-10 shadow-md focus:ring-2 focus:ring-indigo-600 sm:text-sm sm:leading-6
            {error ? 'ring-2 ring-red-600' : ''}"
        aria-describedby="{name}-description"
    >
        <slot />
    </select>
    <p
        id="{name}-description"
        class="inline-block text-xs leading-6
        {error ? 'text-red-600' : 'text-gray-400'}"
    >
        {error || helper}
    </p>
</div>
