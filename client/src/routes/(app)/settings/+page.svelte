<script>
    import Button from "$lib/form/Button.svelte";
    import { toast } from "$lib/overlay/toast";

    /** @type {import("./$types").PageData}*/
    export let data;
    /** @type {import("./$types").ActionData}*/
    export let form;
    $: if (form?.error || data?.error) {
        toast.error("Error", form?.error || data?.error || "Unknown error");
    }
</script>

<h2 class="flex items-center gap-2 text-base/7 font-semibold text-gray-900">
    Profile
</h2>
<p class="mt-1 text-sm/6 text-gray-600">Your account details.</p>
<p class="pt-4 text-sm">
    Email: <span class="font-semibold">{data.user.email}</span>
</p>
<p class="pt-4 text-sm">
    Last activity: <span class="font-semibold">{data.user.updated}</span>
</p>
<hr class="my-8 border-gray-300" />
<div class="grid grid-cols-[1fr_auto] items-center justify-between">
    <div>
        <h2 class="font-semibold leading-7 text-gray-900">Meters</h2>
        <p class="mt-1 text-sm leading-6 text-gray-600">
            List of meters you have access to.
        </p>
    </div>
    <Button href="/settings/meters">Edit Meters</Button>
</div>
<div class="mt-8">
    {#each data.meters as meter}
        <h3 class="mt-4 text-sm font-semibold leading-7 text-gray-900">
            {meter.name}
        </h3>
        <p class="mt-0.5 text-sm leading-6 text-gray-600">
            {meter.description}
        </p>
    {/each}
</div>
