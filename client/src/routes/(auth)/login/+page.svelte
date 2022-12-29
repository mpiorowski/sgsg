<script lang="ts">
    import { goto } from "$app/navigation";
    import { auth } from "$lib/firebase.util";
    import { Button } from "@mpiorowski/svelte-init";
    import { GoogleAuthProvider, signInWithPopup } from "firebase/auth";

    const onLoginWithGoogle = async () => {
        try {
            const response = await signInWithPopup(
                auth,
                new GoogleAuthProvider()
            );
            GoogleAuthProvider.credentialFromResult(response);
            await goto("/"); // redirect to home page
        } catch (error) {
            console.error(error);
        }
    };
</script>

<svelte:head>
    <title>Login</title>
    <meta name="description" content="Login page" />
</svelte:head>

<section
    class="m-auto max-w-xl w-full h-full flex flex-col gap-4 justify-center"
>
    <Button on:click={onLoginWithGoogle}>Login with google</Button>
</section>
