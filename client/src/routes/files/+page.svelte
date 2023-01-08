<script lang="ts">
    import type { PageData } from "./$types";

    export let data: PageData;
</script>

<h1>Files</h1>

{#each data.files as file}
    <p>{file.name}</p>
    <p>{file.id}</p>
    <p>{file.targetId}</p>
{/each}

<form action="?/file" method="post" enctype="multipart/form-data">
    <input type="file" name="file" />
    <input type="submit" value="Upload" />
</form>
