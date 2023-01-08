<script lang="ts">
    import type { PageData } from "./$types";

    export let data: PageData;
</script>

<section>
    <h1>Users</h1>
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
