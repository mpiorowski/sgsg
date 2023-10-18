<script>
    import Input from "$lib/form/Input.svelte";
    import Button from "$lib/form/Button.svelte";
    import { enhance } from "$app/forms";
    import { toast } from "$lib/overlay/toast";
    import SaveIcon from "$lib/icons/SaveIcon.svelte";
    import { extractError } from "$lib/errors";
    import PlusIcon from "$lib/icons/PlusIcon.svelte";
    import Modal from "$lib/overlay/Modal.svelte";
    import DeleteIcon from "$lib/icons/DeleteIcon.svelte";

    /** @type {import("./$types").PageData}*/
    export let data;
    /** @type {import("./$types").ActionData}*/
    export let form;
    $: if (form?.error || data?.error) {
        toast.error("Error", form?.error || data?.error || "Unknown error");
    }

    let name = "";
    let descripton = "";

    let modal = false;
    let id = "";
</script>

{#if modal}
    <Modal
        bind:open={modal}
        title="Are you sure you want to delete this meter?"
        description="This action cannot be undone"
    >
        <form
            action="?/deleteMeter"
            method="post"
            use:enhance={() => {
                return async ({ result, update }) => {
                    if (result.type === "success") {
                        toast.warning("Deleted", "The meter has been deleted");
                    }
                    update();
                    modal = false;
                };
            }}
        >
            <input type="hidden" name="id" bind:value={id} />
            <Button variant="danger">Delete</Button>
        </form>
    </Modal>
{/if}

<form
    action="?/createMeter"
    method="post"
    class="max-w-lg"
    use:enhance={() => {
        return async ({ result, update }) => {
            if (result.type === "success") {
                toast.success("Success", "Meter created");
            }
            update({ reset: false });
        };
    }}
>
    <input type="hidden" name="id" value="" />
    <h2 class="col-span-4 text-lg font-semibold">Add new meter</h2>
    <Input
        bind:value={name}
        name="name"
        label="Name"
        class="mt-8"
        error={extractError(form?.fields, "Name_")}
    />
    <Input
        bind:value={descripton}
        name="description"
        label="Description"
        error={extractError(form?.fields, "Description_")}
    />
    <Button class="col-span-4"><PlusIcon /> Add</Button>
</form>

{#each data.meters as meter}
    <hr class="my-8 border-gray-300" />
    <form
        action="?/createMeter"
        method="post"
        class="max-w-lg"
        use:enhance={() => {
            return async ({ result, update }) => {
                if (result.type === "success") {
                    toast.success("Success", "Meter created");
                }
                update({ reset: false });
            };
        }}
    >
        <input type="hidden" name="id" value={meter.id} />
        <h2 class="mt-6 text-lg font-semibold">{meter.name}</h2>
        <p class="mt-1 text-sm text-gray-600">{meter.description}</p>
        <Input
            bind:value={meter.name}
            name="name"
            label="Name"
            class="mt-4"
            error={extractError(form?.fields, "Name_" + meter.id)}
        />
        <Input
            bind:value={meter.description}
            name="description"
            label="Description"
            error={extractError(form?.fields, "Description_" + meter.id)}
        />
        <div class="flex flex-row gap-2">
            <Button
                class="flex-1"
                variant="danger"
                type="button"
                on:click={() => {
                    id = meter.id;
                    modal = true;
                }}
            >
                <DeleteIcon /> Delete
            </Button>
            <Button class="flex-2"><SaveIcon /> Edit</Button>
        </div>
    </form>
{/each}
