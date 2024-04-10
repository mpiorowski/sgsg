<script>
    import { page } from "$app/stores";

    /** @type {() => void} */
    export let close = () => {};

    $: current = $page.url.pathname.split("/")[1];

    const nav = [
        { name: "profile", title: "Profile" },
        { name: "notes", title: "Notes" },
        { name: "billing", title: "Billing" },
    ];
</script>

<div
    class="flex grow flex-col gap-y-5 overflow-y-auto bg-black/10 px-6 pb-4 ring-1 ring-white/5"
>
    <div class="flex h-16 shrink-0 items-center">
        <a on:click={close} href="/profile" class="flex items-center">
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
        </a>
    </div>
    <nav class="flex flex-1 flex-col">
        <ul role="list" class="flex flex-col items-center space-y-2">
            {#each nav as { name, title }}
                <li class="w-full lg:w-auto">
                    <a
                        on:click={close}
                        href="/{name}"
                        class="group flex rounded-md p-3 text-sm font-semibold leading-6
                    {current === name
                            ? 'bg-gray-800 text-white'
                            : 'text-gray-400 hover:bg-gray-800 hover:text-white'}"
                    >
                        {#if name === "profile"}
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                width="24"
                                height="24"
                                viewBox="0 0 24 24"
                                fill="none"
                                stroke="currentColor"
                                stroke-width="2"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                class="feather feather-user"
                            >
                                <path
                                    d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"
                                ></path>
                                <circle cx="12" cy="7" r="4"></circle>
                            </svg>
                        {:else if name === "notes"}
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                width="24"
                                height="24"
                                viewBox="0 0 24 24"
                                fill="none"
                                stroke="currentColor"
                                stroke-width="2"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                class="feather feather-pen-tool"
                            >
                                <path d="M12 19l7-7 3 3-7 7-3-3z"></path>
                                <path
                                    d="M18 13l-1.5-7.5L2 2l3.5 14.5L13 18l5-5z"
                                ></path>
                                <path d="M2 2l7.586 7.586"></path>
                                <circle cx="11" cy="11" r="2"></circle>
                            </svg>
                        {:else if name === "billing"}
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                width="24"
                                height="24"
                                viewBox="0 0 24 24"
                                fill="none"
                                stroke="currentColor"
                                stroke-width="2"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                class="feather feather-credit-card"
                            >
                                <rect
                                    x="1"
                                    y="4"
                                    width="22"
                                    height="16"
                                    rx="2"
                                    ry="2"
                                ></rect>
                                <line x1="1" y1="10" x2="23" y2="10"></line>
                            </svg>
                        {/if}
                        <span class="lg:sr-only ml-4">{title}</span>
                    </a>
                </li>
            {/each}
        </ul>
    </nav>
</div>
