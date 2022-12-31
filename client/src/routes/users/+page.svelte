<script lang="ts">
    import { Button } from "@mpiorowski/svelte-init";
    import type { PageServerData } from "./$types";

    export let data: PageServerData;
</script>

<section>
    <h2>Users</h2>
    <ul>
        {#each data.users as user}
            <li>
                <pre>
                {JSON.stringify(user, null, 2)}
                </pre>
                <form action="?/delete" method="post" id="delete">
                    <Button form="delete">Delete</Button>
                </form>
            </li>
        {/each}
    </ul>
</section>
