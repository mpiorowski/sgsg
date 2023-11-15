// See https://kit.svelte.dev/docs/types#app

import type { User__Output } from "$lib/proto/proto/User";

// for information about these interfaces
declare global {
    namespace App {
        // interface Error {}
        interface Locals {
            user: User__Output;
            token: string;
        }
        // interface PageData {}
        // interface Platform {}
    }
}

export { };
