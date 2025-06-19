import type { Actions } from './$types';
import { env } from '$env/dynamic/private';

export const actions = {
    createNetwork: async (event) => {
        const formData = await event.request.formData();
        const networkName = formData.get('networkName') as string;
        const token = event.locals.token;

        if (!networkName || !token) {
            return { error: 'Network name and token are required.' };
        }
        try {
            const response = await fetch(`${env.APIBASE}/api/network/create`, {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ name: networkName, domain_id: event.params.domain })
            });
            if (!response.ok) {
                throw new Error(`Failed to create network: ${response.statusText}`);
            }

            const data = await response.json();
            return { success: true, network: data };
        } catch (error) {
            console.error(error);
            return { error: error instanceof Error ? error.message : String(error) };
        }
    },
    deleteNetwork: async (event) => {
        const token = event.locals.token;
        const formData = await event.request.formData();
        const networkId = formData.get('networkId') as string;

        if (!networkId || !token) {
            return { error: 'Domain ID and token are required.' };
        }
        try {
            const response = await fetch(`${env.APIBASE}/api/network/${networkId}`, {
                method: 'DELETE',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ network_id: networkId })
            })
        } catch (error) {
            return { error: error instanceof Error ? error.message : String(error) };
        }
    }
} satisfies Actions;