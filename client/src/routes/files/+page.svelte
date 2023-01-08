<script lang="ts">
    import { enhance } from "$app/forms";
    import { Button } from "@mpiorowski/svelte-init";
    import type { PageData } from "./$types";

    export let data: PageData;
</script>

<h1>Files</h1>

{#each data.files as file}
    <pre>{JSON.stringify(file, null, 2)}</pre>
    <form action="?/deleteFile" method="post" id={file.id} use:enhance>
        <input type="hidden" name="fileId" value={file.id} />
        <input type="hidden" name="targetId" value={file.targetId} />
        <Button type="error" form={file.id}>Delete</Button>
    </form>
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
