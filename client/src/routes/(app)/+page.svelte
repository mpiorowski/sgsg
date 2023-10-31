<script>
    const dl = [
        {
            title: "Efficient and Lightweight Stack",
            description:
                "Utilizes Svelte for the front end, Golang for the back end, and SQLite for a lightweight, self-contained database.",
        },
        {
            title: "High Performance",
            description:
                "Employs gRPC for efficient communication, handling a high volume of concurrent requests with responsiveness.",
        },
        {
            title: "No TypeScript Build Required",
            description:
                "Uses JSDoc comments for type checking, eliminating the need for a TypeScript build step.",
        },
        {
            title: "Minimal External Dependencies",
            description:
                "Reduces security vulnerabilities and simplifies deployment and maintenance by minimizing external dependencies.",
        },
        {
            title: "Efficient Learning and Performance",
            description:
                "Svelte and Golang's simplicity and efficiency make the stack easy to learn, develop, and maintain while delivering high performance.",
        },
    ];
</script>

<div class="flex flex-col items-center justify-center gap-4 p-10 text-gray-900">
    <h1>SGSG</h1>
    <h2>Svelte + Go + SQLite + gRPC</h2>
    <dl class="max-w-xl">
        {#each dl as { title, description }}
            <div class="relative mt-6 pl-9">
                <dt class="font-semibold text-gray-900">
                    <svg
                        class="absolute left-0 top-1 h-5 w-5 text-indigo-500"
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
                    {title}
                </dt>
                <dd class="mt-2 text-sm text-gray-700">
                    {description}
                </dd>
            </div>
        {/each}
    </dl>
</div>
