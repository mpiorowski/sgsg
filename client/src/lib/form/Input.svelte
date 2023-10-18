<script>
    /** @type {string} */
    export let name;
    /** @type {string} */
    export let label;
    /** @type {string | number} */
    export let value;
    /** @type {"text" | "email" | "password"} */
    export let type = "text";
    /** @type {string} */
    export let placeholder = "";
    /** @type {string} */
    export let autocomplete = "off";
    /** @type {number} */
    export let rows = 0;
    /** @type {string} */
    export let error = "";
    /** @type {string} */
    export let helper = "\x80";

    let className = "";
    export { className as class };

    /**
     * @param {HTMLInputElement} node
     * @returns {void}
     */
    function typeAction(node) {
        node.type = type;
    }
</script>

<div class={className}>
    <label for={name} class="block text-sm font-medium leading-6">
        {label}
    </label>
    <div class="mt-2">
        {#if rows === 0}
            <input
                use:typeAction
                bind:value
                id={name}
                {name}
                {placeholder}
                {autocomplete}
                class="block w-full rounded-md border-0 px-3 py-1.5 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6
                {error ? 'ring-2 ring-red-600' : ''}"
                aria-invalid={!!error}
                aria-describedby="{name}-description"
            />
        {:else}
            <textarea
                bind:value
                id={name}
                {name}
                {placeholder}
                {rows}
                class="block w-full rounded-md border-0 px-3 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6
                {error ? 'ring-2 ring-red-600' : ''}"
                aria-invalid={!!error}
                aria-describedby="{name}-description"
            />
        {/if}
    </div>
    <p
        id="{name}-description"
        class="mb-2 mt-0 text-xs leading-6
        {error ? 'text-red-600' : 'text-gray-500'}"
    >
        {error || helper}
    </p>
</div>
