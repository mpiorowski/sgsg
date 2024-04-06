<script>
    import Button from "$lib/form/Button.svelte";
    import Input from "$lib/form/Input.svelte";
    import Checkbox from "$lib/form/Checkbox.svelte";
    import Radio from "$lib/form/Radio.svelte";
    import SelectCustom from "$lib/form/SelectCustom.svelte";
    import Switch from "$lib/form/Switch.svelte";
    import Dropzone from "$lib/form/Dropzone.svelte";
    import FileInput from "$lib/form/FileInput.svelte";
    import Modal from "$lib/ui/Modal.svelte";
    import SelectMultiple from "$lib/form/SelectMultiple.svelte";
    import Tooltip from "$lib/ui/Tooltip.svelte";
    import SelectNative from "$lib/form/SelectNative.svelte";
    import { enhance } from "$app/forms";
    import { showToast, toast } from "$lib/ui/toast.store";
    import { extractError } from "$lib/errors";
    import { generateId } from "$lib/helpers";

    /** @type {import("./$types").PageData} */
    export let data;
    /** @type {import("./$types").ActionData} */
    export let form;
    $: if (form?.error && !form?.fields) {
        toast.error("Error", form.error);
    }
    $: if (form?.fields) {
        const errors = form.fields;
        const firstError = errors[0]?.field.toLowerCase();
        showToast({
            id: generateId(),
            title: "Validation failed",
            description: `Found ${Object.keys(errors).length} errors`,
            type: "error",
            duration: 6000,
            action: {
                label: "Go to first error",
                onClick: () => {
                    /** @type {HTMLInputElement | null} */
                    const input = document.querySelector(
                        `[name="${firstError}"]`,
                    );
                    input?.focus();
                },
            },
        });
    }

    export const countries = /** @type {const} */ ([
        "Poland",
        "United States",
        "Canada",
        "Mexico",
    ]);
    export const skills = /** @type {const} */ ([
        "Frontend Development",
        "Backend Development",
        "UI/UX Design",
        "DevOps",
        "Database Administration",
        "Mobile Development",
    ]);

    /** @type {boolean} */
    let openModal = false;

    let active = data.profile.active ? "on" : "off";
    let email_notifications = data.profile.email_notifications.split(",");
    $: resume_file = new File([data.profile.resume], "resume");
    $: cover_file = new File([data.profile.cover], "cover");
</script>

{#if openModal}
    <Modal
        alert
        bind:open={openModal}
        title="Deactivate your account"
        description="Are you sure you want to deactivate your account? All of your data will be permanently removed. This action cannot be undone."
    >
        <form method="post" action="?/deactive">
            <Button variant="danger">Deactivate</Button>
        </form>
    </Modal>
{/if}

<form
    action="?/update_profile"
    method="post"
    enctype="multipart/form-data"
    class="max-w-2xl"
    use:enhance={() => {
        return async ({ result, update }) => {
            if (result.type === "success") {
                toast.success("Saved", "User profile has been updated.");
            } else {
                // showToast({
                //     id: generateId(),
                //     title: "Validation failed",
                //     description: `Found ${Object.keys(fields).length} errors.`,
                //     type: "error",
                //     duration: 6000,
                //     action: {
                //         label: "Go to first error",
                //         onClick: () => {
                //             /** @type {HTMLInputElement | null} */
                //             const input = document.querySelector(
                //                 `[name="${firstError}"]`,
                //             );
                //             input?.focus();
                //         },
                //     },
                // });
            }
            await update({ reset: false });
        };
    }}
>
    <input type="hidden" name="id" value={data.profile.id} />
    <div class="space-y-12">
        <div class="border-b border-gray-600 pb-12">
            <div>
                <h2
                    class="flex items-center gap-2 text-base font-semibold leading-7"
                >
                    Profile
                    <Tooltip
                        text="This information will be displayed publicly so be careful what you share."
                    >
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="feather feather-info h-4"
                        >
                            <circle cx="12" cy="12" r="10" />
                            <line x1="12" y1="16" x2="12" y2="12" />
                            <line x1="12" y1="8" x2="12.01" y2="8" />
                        </svg>
                    </Tooltip>
                </h2>
                <p class="mt-1 text-sm leading-6 text-gray-400">
                    This information will be displayed publicly so be careful
                    what you share.
                </p>
            </div>

            <div class="mt-10 grid grid-cols-1 gap-x-6 gap-y-2 sm:grid-cols-6">
                <div class="sm:col-span-4">
                    <Switch name="active" label="Active" bind:value={active} />
                </div>
                <div class="sm:col-span-4">
                    <Input
                        name="username"
                        label="Username"
                        autocomplete="username"
                        bind:value={data.profile.username}
                        error={extractError(form?.fields, "username")}
                    />
                </div>

                <div class="col-span-full">
                    <Input
                        name="about"
                        label="About"
                        bind:value={data.profile.about}
                        rows={3}
                        helper="Write a few sentences about yourself."
                        error={extractError(form?.fields, "about")}
                    />
                </div>

                <div class="col-span-full">
                    <FileInput
                        label="Resume"
                        name="resume"
                        bind:file={resume_file}
                        helper="PDF up to 5MB"
                    />
                </div>

                <div class="col-span-full">
                    <Dropzone
                        name="cover"
                        label="Cover photo"
                        bind:file={cover_file}
                        description="SVG, PNG, JPG, GIF up to 10MB"
                        accept="image/*"
                    />
                </div>
            </div>
        </div>

        <div class="border-b border-gray-600 pb-6">
            <h2
                class="flex items-center gap-2 text-base font-semibold leading-7"
            >
                Personal Information

                <Tooltip
                    text="Use a permanent address where you can receive mail."
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        class="feather feather-info h-4"
                    >
                        <circle cx="12" cy="12" r="10" />
                        <line x1="12" y1="16" x2="12" y2="12" />
                        <line x1="12" y1="8" x2="12.01" y2="8" />
                    </svg>
                </Tooltip>
            </h2>
            <p class="mt-1 text-sm leading-6 text-gray-400">
                Use a permanent address where you can receive mail.
            </p>

            <div class="mt-10 grid grid-cols-1 gap-x-6 sm:grid-cols-6">
                <div class="sm:col-span-3">
                    <Input
                        name="first_name"
                        label="First name"
                        autocomplete="given-name"
                        bind:value={data.profile.first_name}
                    />
                </div>

                <div class="sm:col-span-3">
                    <Input
                        name="last_name"
                        label="Last name"
                        autocomplete="family-name"
                        bind:value={data.profile.last_name}
                    />
                </div>

                <div class="sm:col-span-4">
                    <Input
                        name="email"
                        label="Email address"
                        autocomplete="email"
                        bind:value={data.profile.email}
                    />
                </div>

                <div class="sm:col-span-3">
                    <SelectNative
                        name="country"
                        label="Country"
                        bind:value={data.profile.country}
                    >
                        {#each countries as country}
                            <option value={country}>
                                {country}
                            </option>
                        {/each}
                    </SelectNative>
                </div>

                <div class="col-span-full">
                    <Input
                        name="street_address"
                        label="Street address"
                        bind:value={data.profile.street_address}
                        autocomplete="street-address"
                    />
                </div>

                <div class="sm:col-span-2 sm:col-start-1">
                    <Input
                        name="city"
                        label="City"
                        bind:value={data.profile.city}
                        autocomplete="address-level2"
                    />
                </div>

                <div class="sm:col-span-2">
                    <Input
                        name="state"
                        label="State / Province"
                        bind:value={data.profile.state}
                        autocomplete="address-level1"
                    />
                </div>

                <div class="sm:col-span-2">
                    <Input
                        name="zip"
                        label="ZIP / Postal"
                        bind:value={data.profile.zip}
                        autocomplete="postal-code"
                    />
                </div>
            </div>
        </div>

        <div class="border-b border-gray-600 pb-12">
            <h2
                class="flex items-center gap-2 text-base font-semibold leading-7"
            >
                Notifications
                <Tooltip
                    text="We'll always let you know about important changes, but you pick what else you want to hear about."
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        class="feather feather-info h-4"
                    >
                        <circle cx="12" cy="12" r="10" />
                        <line x1="12" y1="16" x2="12" y2="12" />
                        <line x1="12" y1="8" x2="12.01" y2="8" />
                    </svg>
                </Tooltip>
            </h2>
            <p class="mt-1 text-sm leading-6 text-gray-400">
                We'll always let you know about important changes, but you pick
                what else you want to hear about.
            </p>

            <div class="mt-10 space-y-10">
                <fieldset>
                    <legend class="text-sm font-semibold leading-6">
                        By Email
                    </legend>
                    <div class="mt-6 space-y-6">
                        <Checkbox
                            id="comments"
                            name="email_notifications"
                            value="comments"
                            label="Comments"
                            bind:group={email_notifications}
                            description="Get notified when someones posts a comment on a posting."
                        />
                        <Checkbox
                            id="candidates"
                            name="email_notifications"
                            value="candidates"
                            label="Candidates"
                            bind:group={email_notifications}
                            description="Get notified when a candidate applies for a job."
                        />
                        <Checkbox
                            id="offers"
                            name="email_notifications"
                            value="offers"
                            label="Offers"
                            bind:group={email_notifications}
                            description="Get notified when a candidate accepts or rejects an offer."
                        />
                    </div>
                </fieldset>
                <fieldset>
                    <legend class="text-sm font-semibold leading-6">
                        Push Notifications
                    </legend>
                    <p class="mt-1 text-sm leading-6 text-gray-400">
                        These are delivered via SMS to your mobile phone.
                    </p>
                    <div class="mt-6 space-y-6">
                        <Radio
                            id="everything"
                            name="push_notification"
                            label="Everything"
                            value="everything"
                            bind:group={data.profile.push_notification}
                        />
                        <Radio
                            id="mentions"
                            name="push_notification"
                            label="Same as email"
                            value="mentions"
                            bind:group={data.profile.push_notification}
                        />
                        <Radio
                            id="nothing"
                            name="push_notification"
                            label="No push notifications"
                            value="nothing"
                            bind:group={data.profile.push_notification}
                        />
                    </div>
                </fieldset>
            </div>
        </div>
        <div class="pb-12">
            <h2
                class="flex items-center gap-2 text-base font-semibold leading-7"
            >
                Profesional Information
                <Tooltip
                    text="Share your profesional details so others can find you."
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        class="feather feather-info h-4"
                    >
                        <circle cx="12" cy="12" r="10" />
                        <line x1="12" y1="16" x2="12" y2="12" />
                        <line x1="12" y1="8" x2="12.01" y2="8" />
                    </svg>
                </Tooltip>
            </h2>
            <p class="mt-1 text-sm leading-6 text-gray-400">
                Share your profesional details so others can find you.
            </p>
            <div class="mt-6">
                <SelectCustom
                    name="position"
                    label="Position"
                    bind:value={data.profile.position}
                    values={["frontend", "backend"]}
                    options={["Frontend Developer", "Backend Developer"]}
                />
            </div>
            <SelectMultiple
                name="skills"
                label="Skills"
                bind:value={data.profile.skills}
                options={skills}
            />
        </div>
    </div>

    <div class="inline-flex items-center gap-x-4">
        <Button
            type="button"
            variant="link"
            on:click={() => (openModal = true)}
        >
            Deactivate
        </Button>
        <Button>Save</Button>
    </div>
</form>
