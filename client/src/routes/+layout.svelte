<script lang="ts">
    import "./app.css";
</script>

<div class="grid grid-rows-[auto_1fr_auto] h-screen">
    <header class="font-semibold p-2 text-2xl">Go with svelte using grpc</header>

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
