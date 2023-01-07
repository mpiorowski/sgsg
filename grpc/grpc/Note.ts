// Original file: notes.proto

import type { User as _grpc_User, User__Output as _grpc_User__Output } from '../grpc/User';

export interface Note {
  'id'?: (string);
  'created'?: (string);
  'updated'?: (string);
  'deleted'?: (string);
  'userId'?: (string);
  'title'?: (string);
  'content'?: (string);
  'user'?: (_grpc_User | null);
  '_deleted'?: "deleted";
}

export interface Note__Output {
  'id': (string);
  'created': (string);
  'updated': (string);
  'deleted'?: (string);
  'userId': (string);
  'title': (string);
  'content': (string);
  'user': (_grpc_User__Output | null);
  '_deleted': "deleted";
}
