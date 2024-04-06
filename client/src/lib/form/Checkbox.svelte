<script>
    /** @type {string} */
    export let id;
    /** @type {string} */
    export let name;
    /** @type {string} */
    export let label;
    /** @type {string} */
    export let value;
    /** @type {string[]} */
    export let group;
    /** @type {string} */
    export let description = "";

    /** @type {boolean} */
    let checked;

    $: updateChekbox(group);
    $: updateGroup(checked);

    /** @param {string[]} g */
    function updateChekbox(g) {
        checked = g.indexOf(value) >= 0;
    }

    /** @param {boolean} c */
    function updateGroup(c) {
        const index = group.indexOf(value);
        if (c) {
            if (index < 0) {
                group.push(value);
                group = group;
            }
        } else {
            if (index >= 0) {
                group.splice(index, 1);
                group = group;
            }
        }
    }
</script>

<div class="relative flex items-start">
    <div class="flex h-6 items-center">
        <input
            {id}
            {name}
            type="checkbox"
            bind:checked
            bind:value
            class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600"
            aria-describedby="{id}-description"
        />
    </div>
    <div class="ml-3 text-sm leading-6">
        <label for={id} class="font-medium">{label}</label>
        <p id="{id}-description" class="text-gray-400">
            {description}
        </p>
    </div>
</div>
