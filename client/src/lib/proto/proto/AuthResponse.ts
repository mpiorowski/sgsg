// Original file: ../proto/main.proto

import type { User as _proto_User, User__Output as _proto_User__Output } from '../proto/User';

export interface AuthResponse {
  'tokenId'?: (string);
  'user'?: (_proto_User | null);
}

export interface AuthResponse__Output {
  'tokenId': (string);
  'user': (_proto_User__Output | null);
}
