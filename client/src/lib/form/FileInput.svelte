<script>
    /** @type {string} */
    export let name;
    /** @type {string} */
    export let label;
    /** @type {File | undefined} */
    export let file;
    /** @type {string | undefined} */
    export let accept = "*/*";
    /** @type {string | undefined} */
    export let error = "";
    /** @type {string} */
    export let helper = "";

    /** @type {FileList} */
    let files;
    $: if (files && files[0]) {
        file = files[0];
    } else {
        file = undefined;
    }
</script>

<div class="mb-2">
    <label for={name} class="block text-sm font-medium leading-6">
        {label}
    </label>
    <div class="mt-2">
        <input
            bind:files
            {name}
            id={name}
            {accept}
            type="file"
            class="block w-full cursor-pointer rounded-lg border-0 bg-gray-800 shadow-md
            file:mr-3 file:cursor-pointer file:rounded-l-lg file:border-0 file:bg-indigo-600 file:px-3 file:py-1.5 file:text-white file:transition file:hover:bg-indigo-500
            focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-600
            sm:text-sm sm:leading-6
            {error ? 'ring-2 ring-red-600' : ''}"
            aria-invalid={!!error}
            aria-describedby="{name}-description"
        />
    </div>
    <p
        id="{name}-description"
        class="inline-block text-xs leading-6
        {error ? 'text-red-600' : 'text-gray-400'}"
    >
        {error || helper}
    </p>
</div>
