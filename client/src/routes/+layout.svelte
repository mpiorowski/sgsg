<script>
    import "../app.css";
    import Toast from "$lib/ui/toast.svelte";
    import { navigating } from "$app/stores";
    import { toastStore } from "$lib/ui/toast.store";
    import Loading from "./loading.svelte";
</script>

{#if $navigating}
    <Loading />
{/if}

<slot />

<!-- Global notification live region, render this permanently at the end of the document -->
<div
    aria-live="assertive"
    class="pointer-events-none fixed inset-0 z-50 flex items-end px-4 py-6 sm:items-start sm:p-6"
>
    <div class="flex w-full flex-col items-center space-y-4 sm:items-end">
        {#each $toastStore as toast}
            <Toast {toast} />
        {/each}
    </div>
</div>
