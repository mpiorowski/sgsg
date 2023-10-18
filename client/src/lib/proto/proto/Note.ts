// Original file: ../proto/note.proto


export interface Note {
  'id'?: (string);
  'created'?: (string);
  'updated'?: (string);
  'deleted'?: (string);
  'userId'?: (string);
  'title'?: (string);
  'content'?: (string);
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
  '_deleted': "deleted";
}
