import type { ObjectId } from 'mongodb';

export type User = {
    id: ObjectId;
    email: string;
    password_hash: string;
};

export type Domain = {
    id: ObjectId;
    name: string;
    networks: ObjectId[];
    owner: ObjectId;
};

export type VeylNetwork = {
    id: ObjectId;
    name: string;
    relays: ObjectId[];
    resources: ObjectId[];
    domain: ObjectId;
    owner: ObjectId;
};

export type Relay = {
    id: ObjectId;
    public_key: string;
    name: string;
    authentication_key: string;
    network: ObjectId;
};

export type Resource = {
    id: ObjectId;
    name: string;
    alias: string;
    address: string;
    network: ObjectId;
};


