<script lang="ts">
    import { signIn } from "@auth/sveltekit/client";
    import { Button } from "@mpiorowski/svelte-init";
</script>

<svelte:head>
    <title>Login</title>
    <meta name="description" content="Login page" />
</svelte:head>

<section
    class="m-auto max-w-xl w-full h-full flex flex-col gap-4 justify-center"
>
    <Button on:click={() => signIn("google")}>Login with google</Button>
</section>
