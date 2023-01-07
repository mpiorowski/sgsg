// Original file: users.proto

export const UserRole = {
  ROLE_USER: 'ROLE_USER',
  ROLE_ADMIN: 'ROLE_ADMIN',
} as const;

export type UserRole =
  | 'ROLE_USER'
  | 0
  | 'ROLE_ADMIN'
  | 1

export type UserRole__Output = typeof UserRole[keyof typeof UserRole]
