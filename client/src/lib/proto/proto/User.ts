// Original file: proto/auth.proto

import type { Role as _proto_Role, Role__Output as _proto_Role__Output } from '../proto/Role';

export interface User {
  'id'?: (string);
  'created'?: (string);
  'updated'?: (string);
  'deleted'?: (string);
  'email'?: (string);
  'sub'?: (string);
  'role'?: (_proto_Role);
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
  'role': (_proto_Role__Output);
  'avatar': (string);
  'subscription_id': (string);
  'subscription_end': (string);
  'subscription_check': (string);
  'subscription_active': (boolean);
}
