import { GetNetwork } from '$lib/api';

export async function load({ params, locals }) {

    let network = await GetNetwork(locals.token, params.network)

    return {
        network
    };
}