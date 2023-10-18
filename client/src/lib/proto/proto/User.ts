// Original file: ../proto/user.proto


export interface User {
  'id'?: (string);
  'created'?: (string);
  'updated'?: (string);
  'deleted'?: (string);
  'email'?: (string);
  'role'?: (string);
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
  'role': (string);
  'sub': (string);
  'avatar': (string);
  'subscriptionId': (string);
  'subscriptionEnd'?: (string);
  '_deleted': "deleted";
  '_subscriptionEnd': "subscriptionEnd";
}
