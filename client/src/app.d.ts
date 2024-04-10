// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
import type { User__Output } from "$lib/proto/proto/User";
declare global {
    namespace App {
        // interface Error {}
        interface Locals {
            user: User__Output;
            token: string;
        }
        // interface PageData {}
        // interface PageState {}
        // interface Platform {}
    }
}

export { };
