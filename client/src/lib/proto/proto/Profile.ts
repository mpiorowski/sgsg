// Original file: ../proto/profile.proto


export interface Profile {
  'id'?: (string);
  'created'?: (string);
  'updated'?: (string);
  'deleted'?: (string);
  'userId'?: (string);
  'username'?: (string);
  'about'?: (string);
  'resumeId'?: (string);
  'coverId'?: (string);
  'coverUrl'?: (string);
  '_deleted'?: "deleted";
}

export interface Profile__Output {
  'id': (string);
  'created': (string);
  'updated': (string);
  'deleted'?: (string);
  'userId': (string);
  'username': (string);
  'about': (string);
  'resumeId': (string);
  'coverId': (string);
  'coverUrl': (string);
  '_deleted': "deleted";
}
