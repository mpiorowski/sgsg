<script>
    import { onMount } from "svelte";
    import {
        arrow,
        autoPlacement,
        autoUpdate,
        computePosition,
        offset,
        shift,
    } from "@floating-ui/dom";
    import { generateId } from "$lib/helpers";

    /** @type {string} */
    export let text;
    /** @type {boolean} */
    export let auto = false;
    /** @type {HTMLElement} */
    let referenceEl;
    /** @type {HTMLElement} */
    let floatingEl;
    /** @type {HTMLElement} */
    let arrowEl;

    /** @param {HTMLElement} node */
    function portal(node) {
        // add event lsitener for esc to close tooltip
        document.addEventListener("keydown", (e) => {
            if (e.key === "Escape") {
                floatingEl.classList.remove("peer-hover:opacity-100");
                floatingEl.classList.remove("peer-focus:opacity-100");
            }
        });
        node.addEventListener("mouseenter", () => {
            floatingEl.classList.add("peer-hover:opacity-100");
            floatingEl.classList.add("peer-focus:opacity-100");
        });
        node.addEventListener("mouseleave", () => {
            floatingEl.classList.remove("peer-hover:opacity-100");
            floatingEl.classList.remove("peer-focus:opacity-100");
        });
        node.addEventListener("focusin", () => {
            floatingEl.classList.add("peer-hover:opacity-100");
            floatingEl.classList.add("peer-focus:opacity-100");
        });
        return {
            destroy() {
                document.removeEventListener("keydown", (e) => {
                    if (e.key === "Escape") {
                        floatingEl.classList.remove("peer-hover:opacity-100");
                        floatingEl.classList.remove("peer-focus:opacity-100");
                    }
                });
                node.addEventListener("mouseenter", () => {
                    floatingEl.classList.add("peer-hover:opacity-100");
                    floatingEl.classList.add("peer-focus:opacity-100");
                });
                node.addEventListener("mouseleave", () => {
                    floatingEl.classList.remove("peer-hover:opacity-100");
                    floatingEl.classList.remove("peer-focus:opacity-100");
                });
                node.addEventListener("focusin", () => {
                    floatingEl.classList.add("peer-hover:opacity-100");
                    floatingEl.classList.add("peer-focus:opacity-100");
                });
            },
        };
    }

    /**
     * Update the position of the floating element
     */
    function update() {
        computePosition(referenceEl, floatingEl, {
            placement: "top",
            middleware: [
                offset(8),
                autoPlacement(),
                shift({ padding: 10 }),
                arrow({ element: arrowEl }),
            ],
        }).then(({ x, y, placement, middlewareData }) => {
            Object.assign(floatingEl.style, {
                left: `${x}px`,
                top: `${y}px`,
            });

            // Arrow
            if (!middlewareData.arrow) {
                return;
            }
            const arrowPlacement = placement.split("-")[0] ?? "top";
            const { x: arrowX, y: arrowY } = middlewareData.arrow;
            const staticSide =
                {
                    top: "bottom",
                    right: "left",
                    bottom: "top",
                    left: "right",
                }[arrowPlacement] ?? "bottom";

            Object.assign(arrowEl.style, {
                left: arrowX != null ? `${arrowX}px` : "",
                top: arrowY != null ? `${arrowY}px` : "",
                right: "",
                bottom: "",
                [staticSide]: "-4px",
            });
        });
    }

    onMount(() => {
        // for one time update
        if (!auto) {
            update();
            return () => {};
        }
        // for auto update
        else {
            const cleanup = autoUpdate(referenceEl, floatingEl, update);
            return () => cleanup();
        }
    });

    const id = generateId();
</script>

<div
    use:portal
    bind:this={referenceEl}
    aria-describedby="tooltip-{id}"
    class="group inline-flex"
>
    <slot />
    <div
        bind:this={floatingEl}
        id="tooltip-{id}"
        role="tooltip"
        class="invisible absolute w-fit rounded bg-gray-800 p-2 font-normal text-white opacity-0 transition-opacity delay-200
        group-hover:visible group-hover:opacity-100 group-focus:visible group-focus:opacity-100"
    >
        {text}
        <div
            bind:this={arrowEl}
            class="absolute h-2 w-2 rotate-45 bg-gray-800"
        />
    </div>
</div>
