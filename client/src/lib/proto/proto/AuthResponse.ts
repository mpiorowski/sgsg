// Original file: proto/main.proto

import type { User as _proto_User, User__Output as _proto_User__Output } from '../proto/User';

export interface AuthResponse {
  'token'?: (string);
  'user'?: (_proto_User | null);
}

export interface AuthResponse__Output {
  'token': (string);
  'user': (_proto_User__Output | null);
}
