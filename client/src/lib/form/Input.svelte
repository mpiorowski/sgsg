<script>
    /** @type {string} */
    export let name;
    /** @type {string} */
    export let label;
    /** @type {string | number} */
    export let value;
    /** @type {"text" | "email" | "password" | "number" | "date" | "time" | "datetime-local" | "search" | "tel" | "url"} */
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
    export let helper = "";

    /**
     * @param {HTMLInputElement} node
     * @returns {void}
     */
    function typeAction(node) {
        node.type = type;
    }
</script>

<div>
    <label for={name} class="mb-2 block text-sm font-medium leading-6">
        {label}
    </label>
    {#if rows === 0}
        <input
            use:typeAction
            bind:value
            id={name}
            {name}
            {placeholder}
            {autocomplete}
            class="block w-full rounded-md border-0 bg-gray-800 px-3 py-1.5 shadow-md placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6
                {error ? 'outline outline-1 outline-red-600' : ''}"
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
            class="block w-full rounded-md border-0 bg-gray-800 px-3 py-1.5 shadow-md placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6
                {error ? 'outline outline-1 outline-red-600' : ''}"
            aria-invalid={!!error}
            aria-describedby="{name}-description"
        />
    {/if}
    <p
        id="{name}-description"
        class="mb-2 block text-sm leading-6
        {error ? 'text-red-600' : 'text-gray-400'}"
    >
        {error || helper}&#8203;
    </p>
</div>
