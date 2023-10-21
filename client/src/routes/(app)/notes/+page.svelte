<script>
    import { toast } from "$lib/overlay/toast";

    /** @type {import("./$types").PageData}*/
    export let data;
    /** @type {import("./$types").ActionData}*/
    export let form;
    $: if (form?.error || data?.error) {
        toast.error("Error", form?.error || data?.error || "Unknown error");
    }
</script>

<h1>Page</h1>
