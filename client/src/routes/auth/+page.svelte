<script>
    import { page } from "$app/stores";
    import { PUBLIC_SERVER_HTTP } from "$env/static/public";
    import Button from "$lib/form/Button.svelte";
    import { toast } from "$lib/overlay/toast";

    const error = $page.url.searchParams.get("error");
    $: if (error) {
        if (error === "invalid_user") {
            toast.error(
                "Invalid user",
                "User already connected to another provider",
            );
        } else if (error === "unauthorized") {
            toast.error("Unauthorized", "You are not authorized to access");
        } else {
            toast.error("Error", "Something went wrong");
        }
    }

    /**
     * Check if server is running and redirect to login page
     * @param {string} provider
     */
    async function onLogin(provider) {
        try {
            const response = await fetch(`${PUBLIC_SERVER_HTTP}`);
            if (response.status !== 200) {
                toast.error("Error", "Server is not running");
                return;
            }
        } catch (err) {
            console.error(err);
            toast.error("Error", "Server is not running");
            return;
        }
        window.location.href = `${PUBLIC_SERVER_HTTP}/oauth-login/${provider}`;
    }
</script>

<main
    class="flex min-h-full flex-col justify-center bg-gray-900 px-6 py-12 lg:px-8"
>
    <div class="sm:mx-auto sm:w-full sm:max-w-sm">
        <img
            class="mx-auto h-10 w-auto"
            src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=500"
            alt="Your Company"
        />
        <h2
            class="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-white"
        >
            Sign in to your account
        </h2>
    </div>

    <div class="mt-6 sm:mx-auto sm:w-full sm:max-w-sm">
        <div class="mt-6 grid gap-4">
            <Button type="button" on:click={() => onLogin("google")}>
                <svg class="h-5 w-5" viewBox="0 0 24 24">
                    <path
                        d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"
                        fill="white"
                    />
                    <path
                        d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"
                        fill="white"
                    />
                    <path
                        d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"
                        fill="white"
                    />
                    <path
                        d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"
                        fill="white"
                    />
                    <path d="M1 1h22v22H1z" fill="none" />
                </svg>
                <span class="ml-2 text-sm font-semibold leading-6">
                    Continue with Google
                </span>
            </Button>

            <Button on:click={() => onLogin("github")}>
                <svg
                    class="h-5 w-5"
                    fill="currentColor"
                    viewBox="0 0 20 20"
                    aria-hidden="true"
                >
                    <path
                        fill-rule="evenodd"
                        d="M10 0C4.477 0 0 4.484 0 10.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0110 4.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.203 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.942.359.31.678.921.678 1.856 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0020 10.017C20 4.484 15.522 0 10 0z"
                        clip-rule="evenodd"
                    />
                </svg>
                <span class="ml-2 text-sm font-semibold leading-6">
                    Continue with GitHub
                </span>
            </Button>
            <Button on:click={() => onLogin("twitter")}>
                <svg
                    class="h-5 w-5"
                    fill="currentColor"
                    viewBox="0 0 20 20"
                    aria-hidden="true"
                >
                    <path
                        d="M6.29 18.251c7.547 0 11.675-6.253 11.675-11.675 0-.178 0-.355-.012-.53A8.348 8.348 0 0020 3.92a8.19 8.19 0 01-2.357.646 4.118 4.118 0 001.804-2.27 8.224 8.224 0 01-2.605.996 4.107 4.107 0 00-6.993 3.743 11.65 11.65 0 01-8.457-4.287 4.106 4.106 0 001.27 5.477A4.073 4.073 0 01.8 7.713v.052a4.105 4.105 0 003.292 4.022 4.095 4.095 0 01-1.853.07 4.108 4.108 0 003.834 2.85A8.233 8.233 0 010 16.407a11.616 11.616 0 006.29 1.84"
                    />
                </svg>
                <span class="ml-2 text-sm font-semibold leading-6">
                    Continue with Twitter
                </span>
            </Button>
            <!--
            <a
                href="{PUBLIC_SERVER_URL}/oauth-login/google"
                rel="noopener noreferrer"
                class="group inline-flex w-full items-center justify-center gap-x-1.5 rounded-md bg-indigo-500 px-3 py-1.5 text-sm font-semibold text-white shadow-sm transition
                    hover:bg-indigo-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
            >
                <svg class="h-5 w-5" viewBox="0 0 24 24">
                    <path
                        d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"
                        fill="white"
                    />
                    <path
                        d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"
                        fill="white"
                    />
                    <path
                        d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"
                        fill="white"
                    />
                    <path
                        d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"
                        fill="white"
                    />
                    <path d="M1 1h22v22H1z" fill="none" />
                </svg>
                <span class="ml-2 text-sm font-semibold leading-6">
                    Continue with Google
                </span>
            </a>
            <a
                href="{PUBLIC_SERVER_URL}/oauth-login/github"
                rel="noopener noreferrer"
                class="group inline-flex w-full items-center justify-center gap-x-1.5 rounded-md bg-indigo-500 px-3 py-1.5 text-sm font-semibold text-white shadow-sm transition
                    hover:bg-indigo-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
            >
                <svg
                    class="h-5 w-5"
                    fill="currentColor"
                    viewBox="0 0 20 20"
                    aria-hidden="true"
                >
                    <path
                        fill-rule="evenodd"
                        d="M10 0C4.477 0 0 4.484 0 10.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0110 4.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.203 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.942.359.31.678.921.678 1.856 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0020 10.017C20 4.484 15.522 0 10 0z"
                        clip-rule="evenodd"
                    />
                </svg>
                <span class="ml-2 text-sm font-semibold leading-6">
                    Continue with GitHub
                </span>
            </a>
            <a
                href="{PUBLIC_SERVER_URL}/oauth-login/twitter"
                rel="noopener noreferrer"
                class="group inline-flex w-full items-center justify-center gap-x-1.5 rounded-md bg-indigo-500 px-3 py-1.5 text-sm font-semibold text-white shadow-sm transition
                    hover:bg-indigo-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
            >
                <svg
                    class="h-5 w-5"
                    fill="currentColor"
                    viewBox="0 0 20 20"
                    aria-hidden="true"
                >
                    <path
                        d="M6.29 18.251c7.547 0 11.675-6.253 11.675-11.675 0-.178 0-.355-.012-.53A8.348 8.348 0 0020 3.92a8.19 8.19 0 01-2.357.646 4.118 4.118 0 001.804-2.27 8.224 8.224 0 01-2.605.996 4.107 4.107 0 00-6.993 3.743 11.65 11.65 0 01-8.457-4.287 4.106 4.106 0 001.27 5.477A4.073 4.073 0 01.8 7.713v.052a4.105 4.105 0 003.292 4.022 4.095 4.095 0 01-1.853.07 4.108 4.108 0 003.834 2.85A8.233 8.233 0 010 16.407a11.616 11.616 0 006.29 1.84"
                    />
                </svg>
                <span class="ml-2 text-sm font-semibold leading-6">
                    Continue with Twitter
                </span>
            </a>
            -->
        </div>
    </div>
</main>
