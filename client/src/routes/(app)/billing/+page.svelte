<script>
    import { page } from "$app/stores";
    import Button from "$lib/form/Button.svelte";
    import { toast } from "$lib/overlay/toast";

    /** @type {import("./$types").PageData} */
    export let data;
    /** @type {import("./$types").ActionData} */
    export let form;

    $: if (form?.error) {
        toast.error("Error", form.error);
    }

    $: if ($page.url.searchParams.has("success")) {
        toast.success("Success", "Your subscription has been activated.");
    } else if ($page.url.searchParams.has("cancel")) {
        toast.info("Cancelled", "Your subscription has been cancelled.");
    }
</script>

<div class="max-w-7xl px-6 lg:px-8">
    <div class="mx-auto max-w-2xl sm:text-center">
        <h2 class="text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">
            Simple no-tricks pricing
        </h2>
        <p class="mt-6 text-lg leading-8 text-gray-800">
            Distinctio et nulla eum soluta et neque labore quibusdam. Saepe et
            quasi iusto modi velit ut non voluptas in. Explicabo id ut laborum.
        </p>
    </div>
    <div
        class="mx-auto mt-16 max-w-2xl rounded-3xl ring-1 sm:mt-20 lg:mx-0 lg:flex lg:max-w-none
        {data.subscriptionActive ? 'ring-indigo-500' : 'ring-gray-700'}"
    >
        <div class="p-8 sm:p-10 lg:flex-auto">
            <h3 class="text-2xl font-bold tracking-tight text-gray-900">
                Best value for small teams
            </h3>
            <p class="mt-6 text-base leading-7 text-gray-800">
                Lorem ipsum dolor sit amet consect etur adipisicing elit. Itaque
                amet indis perferendis blanditiis repellendus etur quidem
                assumenda.
            </p>
            <div class="mt-10 flex items-center gap-x-4">
                <h4
                    class="flex-none text-sm font-semibold leading-6 text-indigo-500"
                >
                    What’s included
                </h4>
                <div class="h-px flex-auto bg-gray-100"></div>
            </div>
            <ul
                role="list"
                class="mt-8 grid grid-cols-1 gap-4 text-sm leading-6 text-gray-800 sm:grid-cols-2 sm:gap-6"
            >
                <li class="flex gap-x-3">
                    <svg
                        class="h-6 w-5 flex-none text-indigo-500"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                        aria-hidden="true"
                    >
                        <path
                            fill-rule="evenodd"
                            d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
                            clip-rule="evenodd"
                        />
                    </svg>
                    Private forum access
                </li>
                <li class="flex gap-x-3">
                    <svg
                        class="h-6 w-5 flex-none text-indigo-500"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                        aria-hidden="true"
                    >
                        <path
                            fill-rule="evenodd"
                            d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
                            clip-rule="evenodd"
                        />
                    </svg>
                    Member resources
                </li>
                <li class="flex gap-x-3">
                    <svg
                        class="h-6 w-5 flex-none text-indigo-500"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                        aria-hidden="true"
                    >
                        <path
                            fill-rule="evenodd"
                            d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
                            clip-rule="evenodd"
                        />
                    </svg>
                    Entry to annual conference
                </li>
                <li class="flex gap-x-3">
                    <svg
                        class="h-6 w-5 flex-none text-indigo-500"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                        aria-hidden="true"
                    >
                        <path
                            fill-rule="evenodd"
                            d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
                            clip-rule="evenodd"
                        />
                    </svg>
                    Official member t-shirt
                </li>
            </ul>
        </div>
        <div class="-mt-2 p-2 lg:mt-0 lg:w-full lg:max-w-md lg:flex-shrink-0">
            <div
                class="rounded-2xl bg-gray-800 py-10 text-center ring-1 ring-inset ring-gray-900/5 lg:flex lg:flex-col lg:justify-center lg:py-16"
            >
                <div class="mx-auto max-w-xs px-8">
                    <p class="text-base font-semibold text-gray-100">
                        Subscription
                    </p>
                    <p class="mt-6 flex items-baseline justify-center gap-x-2">
                        <span
                            class="text-5xl font-bold tracking-tight text-gray-100"
                        >
                            $349
                        </span>
                        <span
                            class="text-sm font-semibold leading-6 tracking-wide text-gray-100"
                        >
                            /month
                        </span>
                    </p>
                    <div class="mt-8">
                        {#if data.subscriptionActive}
                            <form action="?/createStripePortal" method="post">
                                <Button class="w-full">
                                    Manage subscription
                                </Button>
                            </form>
                        {:else}
                            <form action="?/createStripeCheckout" method="post">
                                <Button class="w-full">Get access</Button>
                            </form>
                        {/if}
                    </div>

                    <p class="mt-6 text-xs leading-5 text-gray-100">
                        Invoices and receipts available for easy company
                        reimbursement
                    </p>
                </div>
            </div>
        </div>
    </div>
</div>
