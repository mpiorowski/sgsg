<script lang="ts">
    import { goto, invalidateAll } from "$app/navigation";
    import { clientAuth } from "$lib/firebase.util";
    import { Button } from "@mpiorowski/svelte-init";
    import { GoogleAuthProvider, signInWithPopup } from "firebase/auth";

    const onLoginWithGoogle = async () => {
        try {
            const google = new GoogleAuthProvider();
            const response = await signInWithPopup(clientAuth, google);
            const idToken = await response.user.getIdToken();
            if (!idToken) {
                throw new Error("No token");
            }
            const form = new FormData();
            form.append("idToken", idToken);
            const result = await fetch("?/login", {
                method: "POST",
                body: form,
            });
            if (!result.ok) {
                throw new Error("Login failed");
            }
            invalidateAll();
            goto("/");
        } catch (error) {
            console.error(error);
            clientAuth.signOut();
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
    <form id="login" method="post" on:submit|preventDefault={onLoginWithGoogle}>
        <Button form="login">Login with google</Button>
    </form>
</section>
