// Original file: users.proto


export interface User {
  'id'?: (string);
  'created'?: (string);
  'updated'?: (string);
  'deleted'?: (string);
  'email'?: (string);
  'role'?: (string);
  'providerId'?: (string);
  '_deleted'?: "deleted";
}

export interface User__Output {
  'id': (string);
  'created': (string);
  'updated': (string);
  'deleted'?: (string);
  'email': (string);
  'role': (string);
  'providerId': (string);
  '_deleted': "deleted";
}
