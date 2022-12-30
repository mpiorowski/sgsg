<script lang="ts">
    import { clientAuth } from "$lib/firebase.client";
    import { inMemoryPersistence } from "firebase/auth";
    import "./app.css";

    clientAuth.setPersistence(inMemoryPersistence)
</script>

<div class="grid grid-rows-[auto_1fr_auto] h-screen">
    <header>
        <div class="font-semibold p-2 text-2xl">Go with svelte using grpc</div>
    </header>

    <main>
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
