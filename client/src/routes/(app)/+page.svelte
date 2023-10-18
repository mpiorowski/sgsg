<script>
    /** @type {import("./$types").PageData}*/
    export let data;
</script>

<div class="text-gray-900">
    <pre class="overflow-x-auto font-mono text-sm">{JSON.stringify(
            data,
            null,
            2,
        )}</pre>
</div>
