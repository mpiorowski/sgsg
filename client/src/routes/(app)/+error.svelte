<script lang="ts">
    import { page } from "$app/stores";
    import Button from "$lib/form/Button.svelte";
</script>

<div class="mt-20 grid place-items-center">
    <div class="text-center">
        <p class="text-base font-semibold text-indigo-600">{$page.status}</p>
        <h1
            class="mt-4 text-3xl font-bold tracking-tight text-gray-50 sm:text-5xl"
        >
            {$page.status === 404 ? "Page not found" : "Something went wrong"}
        </h1>
        <p class="mt-6 text-base leading-7 text-gray-400">
            {$page.error?.message ?? "Unknown error"}
        </p>
        <div class="mt-10 flex items-center justify-center gap-x-6">
            <Button class="flex-1" href="/contact">Go back home</Button>
            <a
                href="/support"
                target="_blank"
                class="flex-1 text-sm font-semibold text-gray-50"
            >
                Contact support
            </a>
        </div>
    </div>
</div>
