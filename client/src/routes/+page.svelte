<script lang="ts">
    import { signOut } from "@auth/sveltekit/client";
    import { Button } from "@mpiorowski/svelte-init";
    import type { PageData } from "./$types";

    export let data: PageData;
</script>

<svelte:head>
    <title>Home</title>
    <meta name="description" content="Home page" />
</svelte:head>

<section class="flex flex-col m-auto max-w-xl gap-2">
    <h1>Home</h1>

    <div>
        <pre>{JSON.stringify(data, undefined, 2)}</pre>
    </div>
        <Button on:click={() => signOut()}>Logout</Button>
</section>
