import type { Actions } from './$types';
import {env} from '$env/dynamic/private';

export const actions = {
	createDomain: async (event) => {
        const formData = await event.request.formData();
        const domainName = formData.get('domainName') as string;
        const token = event.locals.token;

        if (!domainName || !token) {
            return { error: 'Domain name and token are required.' };
        }
        try {
            const response = await fetch(`${env.APIBASE}/api/domain/create`, {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ name: domainName })
            });

            if (!response.ok) {
                throw new Error(`Failed to create domain: ${response.statusText}`);
            }

            const data = await response.json();
            return { success: true, domain: data };
        } catch (error) {
            return { error: error.message };
        }
	}
} satisfies Actions;