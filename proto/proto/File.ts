// Original file: main.proto


export interface File {
  'id'?: (string);
  'created'?: (string);
  'updated'?: (string);
  'deleted'?: (boolean);
  'targetId'?: (string);
  'name'?: (string);
  'type'?: (string);
  'data'?: (Buffer | Uint8Array | string);
  'url'?: (string);
}

export interface File__Output {
  'id': (string);
  'created': (string);
  'updated': (string);
  'deleted': (boolean);
  'targetId': (string);
  'name': (string);
  'type': (string);
  'data': (Buffer);
  'url': (string);
}
