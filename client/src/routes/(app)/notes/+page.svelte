<script>
    import { toast } from "$lib/overlay/toast";
    import Input from "$lib/form/Input.svelte";
    import Button from "$lib/form/Button.svelte";
    import { enhance } from "$app/forms";
    import { extractError } from "$lib/errors";
    import SaveIcon from "$lib/icons/SaveIcon.svelte";

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
    /** @type {boolean} */
    let loading = false;
</script>

<form
    class="m-auto max-w-2xl p-10"
    action="?/insert"
    method="post"
    use:enhance={() => {
        const timeout = setTimeout(() => {
            loading = true;
        }, 100);
        return async ({ result, update }) => {
            if (result.type === "success") {
                toast.success("Success", "Note created");
            }
            clearTimeout(timeout);
            loading = false;
            await update();
        };
    }}
>
    <div class="space-y-12">
        <div>
            <h2
                class="flex items-center gap-2 text-base font-semibold leading-7 text-gray-900"
            >
                Notes
            </h2>
            <p class="mt-1 text-sm leading-6 text-gray-600">
                List of notes you have created.
            </p>
        </div>

        <div class="mt-10 grid grid-cols-1 gap-x-6 gap-y-2 sm:grid-cols-6">
            <div class="sm:col-span-4">
                <Input
                    name="title"
                    label="Title"
                    bind:value={title}
                    error={extractError(form?.fields, "title")}
                />
            </div>

            <div class="col-span-full">
                <Input
                    name="content"
                    label="Content"
                    bind:value={content}
                    error={extractError(form?.fields, "content")}
                    rows={3}
                    helper="Max 1000 characters"
                />
            </div>
            <div class="col-span-full flex justify-end">
                <Button type="submit" {loading}>
                    <svelte:fragment slot="icon">
                        <SaveIcon />
                    </svelte:fragment>
                    Save
                </Button>
            </div>
        </div>

        {#each data.notes as note}
            <div
                class="mx-auto mt-8 rounded-lg bg-gray-700 p-6 text-white shadow-md"
            >
                <h2 class="mb-2 text-xl font-semibold">{note.title}</h2>
                <p class="mb-4 text-gray-50">
                    {note.content}
                </p>
                <Button class="w-20" href="/notes/{note.id}">Edit</Button>
            </div>
        {/each}
    </div>
</form>
