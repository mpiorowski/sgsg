// Original file: main.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { Note as _proto_Note, Note__Output as _proto_Note__Output } from '../proto/Note';
import type { NoteId as _proto_NoteId, NoteId__Output as _proto_NoteId__Output } from '../proto/NoteId';
import type { UserId as _proto_UserId, UserId__Output as _proto_UserId__Output } from '../proto/UserId';

export interface NotesServiceClient extends grpc.Client {
  CreateNote(argument: _proto_Note, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  CreateNote(argument: _proto_Note, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  CreateNote(argument: _proto_Note, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  CreateNote(argument: _proto_Note, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  createNote(argument: _proto_Note, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  createNote(argument: _proto_Note, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  createNote(argument: _proto_Note, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  createNote(argument: _proto_Note, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  
  DeleteNote(argument: _proto_NoteId, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  DeleteNote(argument: _proto_NoteId, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  DeleteNote(argument: _proto_NoteId, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  DeleteNote(argument: _proto_NoteId, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  deleteNote(argument: _proto_NoteId, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  deleteNote(argument: _proto_NoteId, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  deleteNote(argument: _proto_NoteId, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  deleteNote(argument: _proto_NoteId, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  
  GetNotes(argument: _proto_UserId, metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_Note__Output>;
  GetNotes(argument: _proto_UserId, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_Note__Output>;
  getNotes(argument: _proto_UserId, metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_Note__Output>;
  getNotes(argument: _proto_UserId, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_Note__Output>;
  
}

export interface NotesServiceHandlers extends grpc.UntypedServiceImplementation {
  CreateNote: grpc.handleUnaryCall<_proto_Note__Output, _proto_Note>;
  
  DeleteNote: grpc.handleUnaryCall<_proto_NoteId__Output, _proto_Note>;
  
  GetNotes: grpc.handleServerStreamingCall<_proto_UserId__Output, _proto_Note>;
  
}

export interface NotesServiceDefinition extends grpc.ServiceDefinition {
  CreateNote: MethodDefinition<_proto_Note, _proto_Note, _proto_Note__Output, _proto_Note__Output>
  DeleteNote: MethodDefinition<_proto_NoteId, _proto_Note, _proto_NoteId__Output, _proto_Note__Output>
  GetNotes: MethodDefinition<_proto_UserId, _proto_Note, _proto_UserId__Output, _proto_Note__Output>
}
