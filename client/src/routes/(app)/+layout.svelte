<script>
    import Breadcrumbs from "./Breadcrumbs.svelte";
    import Nav from "./Nav.svelte";

    /** @type {import("./$types").PageData} */
    export let data;
</script>

<div id="scroll" class="min-h-full bg-gray-100">
    <div class="bg-gray-800 pb-32">
        <Nav {data} />
        <div class="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
            <Breadcrumbs />
        </div>
    </div>

    <main class="-mt-32">
        <div class="mx-auto max-w-7xl px-4 pb-12 sm:px-6 lg:px-8">
            <div class="rounded-lg bg-white px-5 py-6 shadow sm:px-6">
                <slot />
            </div>
        </div>
    </main>
</div>
