/** @type {import('./$types').LayoutServerLoad} */
export function load({ locals }) {
    return {
        email: locals.user.email,
        avatar: locals.user.avatar,
    };
}
