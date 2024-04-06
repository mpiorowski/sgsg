<script>
    import { toast } from "$lib/ui/toast.store";
    import { enhance } from "$app/forms";
    import Button from "$lib/form/Button.svelte";
    import Input from "$lib/form/Input.svelte";
    import { extractError } from "$lib/errors";
    import Pagination from "$lib/ui/Pagination.svelte";

    /** @type {import("./$types").PageData} */
    export let data;
    /** @type {import("./$types").ActionData} */
    export let form;
    $: if (form?.error) {
        toast.error("Error", form.error);
    }

    /** @type {string} */
    let title = "";
    /** @type {string} */
    let content = "";
    /** @type {boolean} */
    let loading = false;
</script>

<form
    class="max-w-2xl"
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
                class="flex items-center gap-2 text-base font-semibold leading-7"
            >
                Notes
            </h2>
            <p class="mt-1 text-sm leading-6 text-gray-400">
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
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="feather feather-save"
                        >
                            <path
                                d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"
                            ></path>
                            <polyline points="17 21 17 13 7 13 7 21"></polyline>
                            <polyline points="7 3 7 8 15 8"></polyline>
                        </svg>
                    </svelte:fragment>
                    Save
                </Button>
            </div>
        </div>

        {#each data.notes as note}
            <div
                class="mx-auto mt-8 rounded-lg bg-gray-800 p-6 text-white shadow-md"
            >
                <h2 class="mb-2 text-xl font-semibold">{note.title}</h2>
                <p class="mb-4 text-gray-50">
                    {note.content}
                </p>
                <Button class="w-20" href="/notes/{note.id}">Edit</Button>
            </div>
        {/each}
        <Pagination total={data.total} limit={data.limit} />
    </div>
</form>
