<script lang="ts">
    import { UserRole } from "../../../proto/proto/UserRole";
    import type { LayoutData } from "./$types";
    import "./app.css";
    import { navigating } from "$app/stores";
    import { Spinner } from "@mpiorowski/svelte-init";

    export let data: LayoutData;
</script>

<div class="grid grid-rows-[auto_1fr_auto] h-screen">
    <header class="flex flex-row items-center gap-4 font-semibold text-xl p-4">
        <div>Go with svelte using grpc</div>
        <a
            href="/auth"
            class="text-gray-400 hover:text-green-700 border border-gray-400 px-2 rounded"
        >
            Login
        </a>
        <a
            href="/"
            class="text-gray-400 hover:text-green-700 border border-gray-400 px-2 rounded"
        >
            Home
        </a>
        <a
            href="/notes"
            class="text-gray-400 hover:text-green-700 border border-gray-400 px-2 rounded"
        >
            Notes
        </a>
        <a
            href="/email"
            class="text-gray-400 hover:text-green-700 border border-gray-400 px-2 rounded"
        >
            Email
        </a>
        <a
            href="/files"
            class="text-gray-400 hover:text-green-700 border border-gray-400 px-2 rounded"
        >
            Files
        </a>
        {#if data.role === UserRole.ROLE_ADMIN}
            <a
                href="/users"
                class="text-gray-400 hover:text-green-700 border border-gray-400 px-2 rounded"
            >
                Users
            </a>
        {/if}
    </header>

    <main class="max-w-xl m-auto h-full w-full">
        {#if $navigating}
            <div class="absolute top-0 left-0 w-screen h-screen bg-gray-700/30 ">
                <Spinner center />
            </div>
        {/if}
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
