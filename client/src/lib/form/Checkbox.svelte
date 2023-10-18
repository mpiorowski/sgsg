<script>
    /** @type {string} */
    export let name;
    /** @type {string} */
    export let label;
    /** @type {string} */
    export let description = "";
    /** @type {boolean} */
    export let checked = false;
</script>

<div class="relative flex items-start">
    <div class="flex h-6 items-center">
        <input
            type="checkbox"
            bind:checked
            id={name}
            {name}
            class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600"
            aria-describedby="{name}-description"
        />
    </div>
    <div class="ml-3 text-sm leading-6">
        <label for={name} class="font-medium text-gray-900">{label}</label>
        <p id="{name}-description" class="text-gray-500">
            {description}
        </p>
    </div>
</div>
