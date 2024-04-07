<script>
    import { page } from "$app/stores";
    import { env } from "$env/dynamic/public";
    import Button from "$lib/form/Button.svelte";
    import { toast } from "$lib/ui/toast";

    let loading = false;

    const error = $page.url.searchParams.get("error");
    $: if (error) {
        if (error === "invalid_user") {
            toast.error(
                "Invalid user",
                "User already connected to another provider",
            );
        } else if (error === "unauthorized") {
            toast.error("Unauthorized", "You are not authorized to access");
        } else {
            toast.error("Error", "Something went wrong");
        }
    }

    /**
     * Check if server is running and redirect to login page
     * @param {string} provider
     */
    async function onLogin(provider) {
        try {
            loading = true;
            const response = await fetch(`${env.PUBLIC_AUTH_URL}`);
            if (response.status !== 200) {
                toast.error("Error", "Server is not running");
                return;
            }
        } catch (err) {
            console.error(err);
            toast.error("Error", "Server is not running");
            loading = false;
            return;
        }
        window.location.href = `${env.PUBLIC_AUTH_URL}/oauth-login/${provider}`;
    }
</script>

<main
    class="flex min-h-full flex-col justify-center bg-gray-900 px-6 py-12 lg:px-8"
>
    <div class="sm:mx-auto sm:w-full sm:max-w-sm">
        <svg
            version="1.0"
            preserveAspectRatio="xMidYMid meet"
            viewBox="552.22 502.09 972.78 1036.54"
            class="mx-auto h-10 w-10 text-indigo-600"
        >
            <g
                transform="translate(0.000000,2048.000000) scale(0.100000,-0.100000)"
                fill="currentColor"
                stroke="none"
            >
                <path
                    d="M10040 15453 c-264 -43 -362 -62 -523 -102 -358 -90 -617 -186 -952 -350 -305 -150 -537 -296 -791 -498 -140 -111 -514 -486 -628 -628 -408 -512 -635 -1104 -653 -1700 -9 -298 23 -530 107 -785 206 -623 722 -1122 1259 -1215 70 -13 252 -15 1089 -15 l1006 0 -89 51 c-137 78 -364 221 -520 327 -277 188 -463 341 -665 549 -209 214 -339 392 -450 618 -316 642 -303 1237 40 1926 260 522 674 992 1255 1425 171 128 362 257 514 348 85 51 91 56 60 55 -19 -1 -45 -4 -59 -6z"
                />
                <path
                    d="M10830 14680 c-674 -47 -1346 -309 -1805 -702 -387 -332 -610 -708 -675 -1138 -16 -110 -13 -356 6 -469 44 -261 150 -502 317 -726 82 -109 338 -362 472 -467 277 -216 310 -233 1755 -951 765 -380 877 -440 1124 -604 457 -303 746 -613 930 -994 123 -254 162 -432 153 -694 -8 -233 -50 -395 -171 -662 -180 -399 -472 -799 -891 -1219 -340 -342 -552 -516 -1048 -861 -104 -72 -142 -103 -122 -99 17 3 87 14 156 25 751 120 1504 436 2175 916 234 168 395 306 624 535 293 292 479 518 675 820 325 500 569 1083 674 1610 35 174 71 443 71 525 l0 55 -758 0 -758 0 -39 88 c-124 275 -418 565 -814 802 -142 85 -640 335 -1416 710 -872 421 -1006 494 -1119 606 -280 275 -201 594 205 833 239 141 499 201 867 201 331 0 588 -42 980 -161 521 -158 1090 -445 1557 -784 454 -330 931 -817 1205 -1230 44 -66 80 -113 80 -105 0 35 -43 306 -71 445 -218 1100 -777 2062 -1604 2756 -607 510 -1364 841 -2110 924 -169 19 -467 26 -625 15z"
                />
                <path
                    d="M5525 10340 c2 -14 7 -54 10 -90 29 -314 167 -857 305 -1205 242 -607 556 -1095 994 -1543 664 -680 1539 -1114 2491 -1239 175 -23 763 -26 940 -5 823 98 1520 386 2033 842 200 177 353 399 435 627 85 236 72 477 -37 708 -137 289 -379 540 -736 763 -124 77 -430 236 -430 223 0 -5 5 -13 11 -17 6 -3 26 -38 45 -77 253 -529 -272 -1160 -1161 -1392 -269 -70 -474 -95 -786 -95 -134 0 -286 6 -374 15 -693 73 -1446 361 -2120 810 -382 256 -785 612 -1084 960 -130 151 -334 423 -437 582 -115 177 -104 162 -99 133z"
                />
            </g>
        </svg>
        <h2
            class="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-white"
        >
            Sign in to your account
        </h2>
    </div>

    <div class="mt-6 sm:mx-auto sm:w-full sm:max-w-sm">
        <div class="mt-6 grid gap-4">
            <Button type="button" on:click={() => onLogin("google")} {loading}>
                <svelte:fragment slot="icon">
                    <svg class="h-5 w-5" viewBox="0 0 24 24">
                        <path
                            d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"
                            fill="white"
                        />
                        <path
                            d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"
                            fill="white"
                        />
                        <path
                            d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"
                            fill="white"
                        />
                        <path
                            d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"
                            fill="white"
                        />
                        <path d="M1 1h22v22H1z" fill="none" />
                    </svg>
                </svelte:fragment>
                <span class="ml-2 text-sm font-semibold leading-6">
                    Continue with Google
                </span>
            </Button>

            <Button on:click={() => onLogin("github")} {loading}>
                <svelte:fragment slot="icon">
                    <svg
                        class="h-5 w-5"
                        fill="currentColor"
                        viewBox="0 0 20 20"
                        aria-hidden="true"
                    >
                        <path
                            fill-rule="evenodd"
                            d="M10 0C4.477 0 0 4.484 0 10.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0110 4.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.203 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.942.359.31.678.921.678 1.856 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0020 10.017C20 4.484 15.522 0 10 0z"
                            clip-rule="evenodd"
                        />
                    </svg>
                </svelte:fragment>
                <span class="ml-2 text-sm font-semibold leading-6">
                    Continue with GitHub
                </span>
            </Button>
        </div>
    </div>
</main>
