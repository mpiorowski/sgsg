<script>
    import Button from "$lib/form/Button.svelte";
    import Dropzone from "$lib/form/Dropzone.svelte";
    import FileInput from "$lib/form/FileInput.svelte";
    import Input from "$lib/form/Input.svelte";
    import SaveIcon from "$lib/icons/SaveIcon.svelte";

    /** @type {import("./$types").PageData} */
    export let data;

    /** @type {import("./$types").ActionData} */
    // export let form;

    /** @type {File} */
    let resume = new File([""], "resume.pdf", { type: "application/pdf" });

    /** @type {File} */
    let cover = new File([""], "cover.png", { type: "image/png" });
</script>

<div class="text-gray-900">
    <form class="m-auto max-w-2xl p-10">
        <div class="space-y-12">
            <div>
                <h2
                    class="flex items-center gap-2 text-base font-semibold leading-7 text-gray-900"
                >
                    Profile
                </h2>
                <p class="mt-1 text-sm leading-6 text-gray-600">
                    Your public profile is how other people will see you.
                </p>
            </div>

            <div class="mt-10 grid grid-cols-1 gap-x-6 gap-y-2 sm:grid-cols-6">
                <div class="sm:col-span-4">
                    <Input
                        name="username"
                        label="Username"
                        autocomplete="username"
                        bind:value={data.profile.username}
                    />
                </div>

                <div class="col-span-full">
                    <Input
                        name="about"
                        label="About"
                        bind:value={data.profile.about}
                        rows={3}
                        helper="Write a few sentences about yourself."
                    />
                </div>

                <div class="col-span-full">
                    <FileInput
                        name="resume"
                        label="Resume"
                        bind:file={resume}
                        helper="PDF up to 5MB"
                    />
                </div>

                <div class="col-span-full mt-6">
                    <Dropzone
                        name="cover"
                        label="Cover photo"
                        bind:file={cover}
                        description="SVG, PNG, JPG, GIF up to 10MB"
                        accept="image/*"
                    />
                </div>
            </div>
            <div class="flex justify-end">
                <Button type="submit">
                    <SaveIcon />
                    Save
                </Button>
            </div>
        </div>
    </form>
</div>
