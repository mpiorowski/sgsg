// Original file: main.proto


export interface File {
  'id'?: (string);
  'created'?: (string);
  'updated'?: (string);
  'deleted'?: (string);
  'targetId'?: (string);
  'name'?: (string);
  'type'?: (string);
  'data'?: (Buffer | Uint8Array | string);
  'url'?: (string);
  '_deleted'?: "deleted";
}

export interface File__Output {
  'id': (string);
  'created': (string);
  'updated': (string);
  'deleted'?: (string);
  'targetId': (string);
  'name': (string);
  'type': (string);
  'data': (Buffer);
  'url': (string);
  '_deleted': "deleted";
}
