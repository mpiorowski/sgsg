<script lang="ts">
    import { enhance } from "$app/forms";
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
                <form action="?/delete" use:enhance method="post" id={user.id}>
                    <input type="hidden" name="user" value={JSON.stringify(user)} />
                    <Button form={user.id}>Delete</Button>
                </form>
            </li>
        {/each}
    </ul>
</section>
