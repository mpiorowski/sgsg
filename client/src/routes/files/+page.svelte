<script lang="ts">
    import { enhance } from "$app/forms";
    import { Button } from "@mpiorowski/svelte-init";
    import type { PageData } from "./$types";

    export let data: PageData;
</script>

<h1>Files</h1>

{#each data.files as file}
    <pre>{JSON.stringify(file, null, 2)}</pre>
    <Button
        on:click={() => {
            window.open(file.url, "_blank");
        }}
    >
        Download
    </Button>
{/each}

<form action="?/createFile" method="post" id="file" use:enhance>
    <input type="file" name="file" />
    <Button form="file">Upload</Button>
</form>
