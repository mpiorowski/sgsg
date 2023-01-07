export type User = {
    id: string;
    created: Date;
    updated: Date;
    deleted: Date;

    email: string;
    role: Role;
    providerId: string;
}

export enum Role {
    Admin = 'admin',
    User = 'user',
}
