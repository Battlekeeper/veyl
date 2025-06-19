import type { Domain, User, VeylNetwork } from "./types";
import { env } from "$env/dynamic/private"

export async function GetUser(token: string) {
    let res = await fetch(`${env.APIBASE}/api/user`, {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
        }
    })
    if (!res.ok) {
        throw new Error(`Failed to fetch user: ${res.statusText}`);
    }
    let data = await res.json()
    return data as Promise<User>;
}

export async function GetDomains(token: string) {
    let res = await fetch(`${env.APIBASE}/api/user/domains`, {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
        }
    })
    if (!res.ok) {
        throw new Error(`Failed to fetch domain: ${res.statusText}`);
    }
    let data = await res.json()
    return data as Promise<Domain[]>;
}

export async function GetDomain(token: string, domainId: string) {
    let res = await fetch(`${env.APIBASE}/api/domain/${domainId}`, {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
        }
    })
    if (!res.ok) {
        throw new Error(`Failed to fetch domain: ${res.statusText}`);
    }
    let data = await res.json()
    return data as Promise<Domain>;
}

export async function GetDomainNetworks(token: string, domainId: string) {
    let res = await fetch(`${env.APIBASE}/api/domain/${domainId}/networks`, {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
        }
    })
    if (!res.ok) {
        throw new Error(`Failed to fetch domain: ${res.statusText}`);
    }
    let data = await res.json()
    return data as Promise<VeylNetwork[]>;
}

export async function CreateDomain(token: string, name: string) {
    let res = await fetch(`${env.APIBASE}/api/domain/create`, {
        method: 'POST',
        headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ name })
    })
    if (!res.ok) {
        throw new Error(`Failed to create domain: ${res.statusText}`);
    }
    let data = await res.json()
    return data as Promise<Domain>;
}

export async function GetNetwork(token: string, networkId: string) {
    let res = await fetch(`${env.APIBASE}/api/network/${networkId}`, {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
        }
    })
    if (!res.ok) {
        throw new Error(`Failed to fetch network: ${res.statusText}`);
    }
    let data = await res.json()
    return data as Promise<VeylNetwork>;
}