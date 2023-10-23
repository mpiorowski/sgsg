<script>
    import { toast } from "$lib/overlay/toast";
    import Input from "$lib/form/Input.svelte";
    import Button from "$lib/form/Button.svelte";
    import { enhance } from "$app/forms";
    import { extractError } from "$lib/errors";

    /** @type {import("./$types").PageData} */
    export let data;
    /** @type {import("./$types").ActionData} */
    export let form;
    $: if (form?.error || data?.error) {
        toast.error("Error", form?.error || data?.error || "Unknown error");
    }

    /** @type {string} */
    let title = "";
    /** @type {string} */
    let content = "";
</script>

<h1>Notes</h1>

<form
    action="?/insert"
    method="post"
    use:enhance={() => {
        return async ({ result, update }) => {
            if (result.type === "success") {
                toast.success("Success", "Note created");
            }
            await update();
        };
    }}
>
    <Input
        name="title"
        label="Title"
        bind:value={title}
        error={extractError(form?.fields, "title")}
    />
    <Input
        name="content"
        label="Content"
        bind:value={content}
        error={extractError(form?.fields, "content")}
    />
    <Button class="w-20">Create</Button>
</form>

{#each data.notes as note}
    <div class="mt-4 flex gap-4 rounded flex-col border p-2">
        <h2 class="text-xl font-bold">{note.title}</h2>
        <p>{note.content}</p>
        <Button class="w-20" href="/notes/{note.id}">Edit</Button>
    </div>
{/each}
