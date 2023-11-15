// Original file: ../proto/user.proto

export const UserRole = {
  ROLE_UNSPECIFIED: 0,
  ROLE_USER: 1,
  ROLE_ADMIN: 2,
} as const;

export type UserRole =
  | 'ROLE_UNSPECIFIED'
  | 0
  | 'ROLE_USER'
  | 1
  | 'ROLE_ADMIN'
  | 2

export type UserRole__Output = typeof UserRole[keyof typeof UserRole]
