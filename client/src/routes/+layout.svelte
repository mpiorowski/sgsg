<script>
    import Toast from "$lib/overlay/Toast.svelte";
    import { toastStore } from "$lib/overlay/toast";
    import "../app.css";
</script>

<slot />

<!-- Global notification live region, render this permanently at the end of the document -->
<div
    aria-live="assertive"
    class="pointer-events-none fixed inset-0 z-20 flex items-end px-4 py-6 sm:items-start sm:p-6"
>
    <div class="flex w-full flex-col items-center space-y-4 sm:items-end">
        {#each $toastStore as toast}
            <Toast {toast} />
        {/each}
    </div>
</div>
