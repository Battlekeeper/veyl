import {GetDomains} from '$lib/api';

export const load = async ({ params, cookies, locals }) => {
    let domains = await GetDomains(locals.token);

    return {
        domains: domains,
    }
}