// Original file: proto/auth.proto

export const Role = {
  ROLE_UNSET: 0,
  ROLE_USER: 1,
  ROLE_ADMIN: 2,
} as const;

export type Role =
  | 'ROLE_UNSET'
  | 0
  | 'ROLE_USER'
  | 1
  | 'ROLE_ADMIN'
  | 2

export type Role__Output = typeof Role[keyof typeof Role]
