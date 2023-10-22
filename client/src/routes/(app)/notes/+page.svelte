<script>
    import { toast } from "$lib/overlay/toast";
    import Input from "$lib/form/Input.svelte";
    import Button from "$lib/form/Button.svelte";
    import { enhance } from "$app/forms";

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
    <Input name="title" label="Title" bind:value={title} />
    <Input name="content" label="Content" bind:value={content} />
    <Button>Create</Button>
</form>

{#each data.notes as note}
    <div>
        <h2>{note.title}</h2>
        <p>{note.content}</p>
    </div>
{/each}
