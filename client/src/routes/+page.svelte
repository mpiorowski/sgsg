<script lang="ts">
    import { Button } from "@mpiorowski/svelte-init";
    import type { PageServerData } from "./$types";
    export let data: PageServerData;
</script>

<svelte:head>
    <title>Home</title>
    <meta name="description" content="Home page" />
</svelte:head>

<section class="flex flex-col m-auto max-w-xl gap-2">
    <h1>Home</h1>
    <div>Id: {data.user.id}</div>

    <div>
        Email: {data.user.email}
    </div>
    <div>
        Roles: {data.user.role}
    </div>
    <div>
        ProviderId: {data.user.providerId}
    </div>
    <form action="/login?/logout" id="logout" method="post">
        <Button form="logout">Logout</Button>
    </form>
</section>
