export type User = {
    id: string;
    email: string;
    role: UserRole;
    providerId: string;
}

enum UserRole {
    ROLE_USER = "ROLE_USER",
    ROLE_ADMIN = "ROLE_ADMIN",
}
