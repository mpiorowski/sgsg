<script lang="ts">
    import { enhance } from "$app/forms";
    import { Button, Input } from "@mpiorowski/svelte-init";

    let email = "";
    let message = "";
</script>

<h1>Email</h1>

<form action="?/sendEmail" id="sendEmail" method="post" use:enhance>
    <Input bind:value={email} name="email" label="Email" />
    <Input bind:value={message} name="message" label="Message" />
    <Button form="sendEmail">Send</Button>
</form>
