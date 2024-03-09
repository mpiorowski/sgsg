// Original file: proto/auth.proto

import type { UserRole as _proto_UserRole, UserRole__Output as _proto_UserRole__Output } from '../proto/UserRole';

export interface User {
  'id'?: (string);
  'created'?: (string);
  'updated'?: (string);
  'deleted'?: (string);
  'email'?: (string);
  'sub'?: (string);
  'role'?: (_proto_UserRole);
  'avatar'?: (string);
  'subscription_id'?: (string);
  'subscription_end'?: (string);
  'subscription_check'?: (string);
  'subscription_active'?: (boolean);
}

export interface User__Output {
  'id': (string);
  'created': (string);
  'updated': (string);
  'deleted': (string);
  'email': (string);
  'sub': (string);
  'role': (_proto_UserRole__Output);
  'avatar': (string);
  'subscription_id': (string);
  'subscription_end': (string);
  'subscription_check': (string);
  'subscription_active': (boolean);
}
