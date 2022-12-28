<script lang="ts">
    import { Button } from "@mpiorowski/svelte-init";
</script>

<svelte:head>
    <title>Login</title>
    <meta name="description" content="Login page" />
</svelte:head>

<section
    class="m-auto max-w-xl w-full h-full flex flex-col gap-4 justify-center"
>
    <form action="?/loginGoogle" method="post">
        <Button>Login with google</Button>
    </form>
    <form action="?/loginMicrosoft" method="post">
        <Button>Login with microsoft</Button>
    </form>
</section>
