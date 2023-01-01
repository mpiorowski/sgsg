<script lang="ts">
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
            </li>
        {/each}
    </ul>
</section>
