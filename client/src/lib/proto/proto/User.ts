// Original file: ../proto/user.proto

import type { UserRole as _proto_UserRole, UserRole__Output as _proto_UserRole__Output } from '../proto/UserRole';

export interface User {
  'id'?: (string);
  'created'?: (string);
  'updated'?: (string);
  'deleted'?: (string);
  'email'?: (string);
  'role'?: (_proto_UserRole);
  'sub'?: (string);
  'avatar'?: (string);
  'subscriptionId'?: (string);
  'subscriptionEnd'?: (string);
  '_deleted'?: "deleted";
  '_subscriptionEnd'?: "subscriptionEnd";
}

export interface User__Output {
  'id': (string);
  'created': (string);
  'updated': (string);
  'deleted'?: (string);
  'email': (string);
  'role': (_proto_UserRole__Output);
  'sub': (string);
  'avatar': (string);
  'subscriptionId': (string);
  'subscriptionEnd'?: (string);
  '_deleted': "deleted";
  '_subscriptionEnd': "subscriptionEnd";
}
