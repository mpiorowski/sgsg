import type { PageServerLoad, Actions } from './$types';

export const load = (({ cookies }) => {

    return {};
}) satisfies PageServerLoad;

export const actions = {
    loginGoogle: async ({ request, cookies }) => {

    },
    loginMicrosoft: async ({ request, cookies }) => {

    },
} satisfies Actions;
