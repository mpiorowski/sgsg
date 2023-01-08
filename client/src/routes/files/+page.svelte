<script lang="ts">
    import { enhance } from "$app/forms";
    import type { PageData } from "./$types";

    export let data: PageData;
</script>

<h1>Files</h1>

{#each data.files as file}
    <p>{file.name}</p>
    <p>{file.id}</p>
    <p>{file.targetId}</p>
    <a href={file.url}>Download</a>
{/each}

<form action="?/createFile" method="post" id="file" use:enhance>
    <input type="file" name="file" />
    <button type="submit" form="file">Upload</button>
</form>
