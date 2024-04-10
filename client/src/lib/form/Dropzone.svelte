<script>
    /** @type {string} */
    export let name;
    /** @type {string} */
    export let label;
    /** @type {File | undefined} */
    export let file;
    /** @type {string | undefined} */
    export let description;
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

    /** @type {string} */
    export let url = "";
    $: if (file && file?.size > 0) {
        url = URL.createObjectURL(file);
    }
</script>

<div class="mb-2">
    <label for={name} class="block text-sm font-medium leading-6">
        {label}
    </label>
    <label
        id={label}
        for={name}
        class="mt-2 flex h-[200px] cursor-pointer justify-center rounded-lg border border-dashed border-gray-600 bg-gray-800
        focus-within:border-0 focus-within:ring-2 focus-within:ring-indigo-600 hover:border-gray-500
        {error ? 'border-0 ring-2 ring-red-600' : ''}"
    >
        <input
            bind:files
            id={name}
            {name}
            {accept}
            type="file"
            class="sr-only"
        />
        {#if url}
            <img
                class="w-full rounded object-cover text-gray-300"
                src={url}
                alt={name}
            />
        {:else}
            <div class="flex flex-col justify-center text-center">
                <svg
                    class="mx-auto h-12 w-12 text-gray-300"
                    viewBox="0 0 24 24"
                    fill="currentColor"
                    aria-hidden="true"
                >
                    <path
                        fill-rule="evenodd"
                        d="M1.5 6a2.25 2.25 0 012.25-2.25h16.5A2.25 2.25 0 0122.5 6v12a2.25 2.25 0 01-2.25 2.25H3.75A2.25 2.25 0 011.5 18V6zM3 16.06V18c0 .414.336.75.75.75h16.5A.75.75 0 0021 18v-1.94l-2.69-2.689a1.5 1.5 0 00-2.12 0l-.88.879.97.97a.75.75 0 11-1.06 1.06l-5.16-5.159a1.5 1.5 0 00-2.12 0L3 16.061zm10.125-7.81a1.125 1.125 0 112.25 0 1.125 1.125 0 01-2.25 0z"
                        clip-rule="evenodd"
                    />
                </svg>
                {#if file && file.size > 0}
                    <p class="mt-4 flex text-sm leading-6 text-gray-900">
                        {file.name}
                    </p>
                    <p class="text-xs leading-5 text-gray-400">
                        {file.size} bytes
                    </p>
                {:else}
                    <div class="mt-4 flex text-sm leading-6">
                        <span class="font-semibold">
                            Click to upload an image
                        </span>
                    </div>
                    <p class="text-xs leading-5 text-gray-400">
                        {description}
                    </p>
                {/if}
            </div>
        {/if}
    </label>
    <p
        id="{name}-description"
        class="inline-block text-xs leading-6
        {error ? 'text-red-600' : 'text-gray-400'}"
    >
        {error || helper}
    </p>
</div>
