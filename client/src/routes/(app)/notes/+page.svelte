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
    action="?/create"
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
    <Button>Create</Button>
</form>

{#each data.notes as note}
    <div class="mt-4 rounded border p-2">
        <form
            action="?/create"
            method="post"
            use:enhance={() => {
                return async ({ result, update }) => {
                    if (result.type === "success") {
                        toast.success("Success", "Note updated");
                    }
                    await update({ reset: false });
                };
            }}
        >
            <input type="hidden" name="id" value={note.id} />
            <Input
                name="title"
                label="Title"
                bind:value={note.title}
                error={extractError(form?.fields, "title")}
            />
            <Input
                name="content"
                label="Content"
                bind:value={note.content}
                error={extractError(form?.fields, "content")}
            />
            <Button>Update</Button>
        </form>
        <form
            action="?/delete"
            method="post"
            use:enhance={() => {
                return async ({ result, update }) => {
                    if (result.type === "success") {
                        toast.warning("Success", "Note deleted");
                    }
                    await update();
                };
            }}
        >
            <input type="hidden" name="id" value={note.id} />
            <Button variant="danger">Delete</Button>
        </form>
    </div>
{/each}
