<script>
    import { enhance } from "$app/forms";
    import { goto } from "$app/navigation";
    import { extractError } from "$lib/errors";
    import Button from "$lib/form/Button.svelte";
    import Input from "$lib/form/Input.svelte";
    import { toast } from "$lib/overlay/toast";

    /** @type {import("./$types").PageData} */
    export let data;

    /** @type {import("./$types").ActionData} */
    export let form;
    $: if (form?.error) {
        toast.error("Error", form.error);
    }
</script>

<form
    action="?/update"
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
    <input type="hidden" name="id" value={data.note.id} />
    <Input
        name="title"
        label="Title"
        bind:value={data.note.title}
        error={extractError(form?.fields, "title")}
    />
    <Input
        name="content"
        label="Content"
        bind:value={data.note.content}
        error={extractError(form?.fields, "content")}
    />
    <Button class="w-20">Update</Button>
</form>
<form
    action="?/delete"
    method="post"
    use:enhance={() => {
        return async ({ result }) => {
            if (result.type === "success") {
                toast.warning("Success", "Note deleted");
                await goto("/notes");
            }
        };
    }}
>
    <input type="hidden" name="id" value={data.note.id} />
    <Button class="mt-4 w-20" variant="danger">Delete</Button>
</form>
