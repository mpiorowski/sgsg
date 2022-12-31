<script lang="ts">
    import { clientAuth } from "$lib/firebase.client";
    import { inMemoryPersistence } from "firebase/auth";
    import "./app.css";

    clientAuth.setPersistence(inMemoryPersistence)
</script>

<div class="grid grid-rows-[auto_1fr_auto] h-screen">
    <header class="flex flex-row items-center gap-4 font-semibold text-xl p-4">
        <div>Go with svelte using grpc</div>
        <a href="/login" class="text-gray-400 hover:text-green-700 border border-gray-400 px-2 rounded">Login</a>
        <a href="/" class="text-gray-400 hover:text-green-700 border border-gray-400 px-2 rounded">Home</a>
        <a href="/users" class="text-gray-400 hover:text-green-700 border border-gray-400 px-2 rounded">Users</a>
    </header>

    <main class="max-w-xl m-auto h-full w-full">
        <slot />
    </main>

    <footer class="p-2">
        <a
            href="https://www.github.com/mpiorowski/go-svelte-grpc"
            target="_blank"
            rel="noopener noreferrer"
        >
            <img
                src="https://www.vectorlogo.zone/logos/github/github-ar21.svg"
                alt="Github logo"
            />
        </a>
    </footer>
</div>
