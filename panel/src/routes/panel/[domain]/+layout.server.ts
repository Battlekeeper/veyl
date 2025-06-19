import { GetDomain, GetDomainNetworks } from '$lib/api';

export async function load({ params, locals }) {

    let domain = await GetDomain(locals.token, params.domain);
    let networks = await GetDomainNetworks(locals.token, params.domain);

    return {
        networks,
        domain
    };
}