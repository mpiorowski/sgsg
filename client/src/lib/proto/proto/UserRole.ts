// Original file: proto/auth.proto

export const UserRole = {
  ROLE_UNSET: 'ROLE_UNSET',
  ROLE_USER: 'ROLE_USER',
  ROLE_ADMIN: 'ROLE_ADMIN',
} as const;

export type UserRole =
  | 'ROLE_UNSET'
  | 0
  | 'ROLE_USER'
  | 1
  | 'ROLE_ADMIN'
  | 2

export type UserRole__Output = typeof UserRole[keyof typeof UserRole]
