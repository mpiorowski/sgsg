<script lang="ts">
    import { enhance } from "$app/forms";
    import { Button, Input } from "@mpiorowski/svelte-init";
    import type { Note } from "src/types/note.type";
    import type { PageData } from "./$types";

    export let data: PageData;
    let note: Note = {
        id: "",
        userId: "",
        title: "",
        content: "",
    };
</script>

<section>
    <h1>Notes</h1>
    <div>
        for user {data.session?.user?.email}
    </div>
    <form action="?/create" method="post" id="create" use:enhance>
        <Input
            name="title"
            label="Title"
            placeholder="Title"
            value={note.title}
        />
        <Input
            name="content"
            label="Content"
            placeholder="Content"
            value={note.content}
        />
        <Button form="create">Create</Button>
    </form>
    <ul>
        {#each data.notes as note}
            <li>
                <pre>{JSON.stringify(note, undefined, 2)}</pre>
                <form action="?/delete" method="post" use:enhance id={note.id}>
                    <input type="hidden" name="id" value={note.id} />
                    <Button type="error" form={note.id}>Delete</Button>
                </form>
            </li>
        {/each}
    </ul>
</section>
