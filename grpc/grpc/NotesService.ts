// Original file: grpc.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { Note as _grpc_Note, Note__Output as _grpc_Note__Output } from '../grpc/Note';
import type { NoteId as _grpc_NoteId, NoteId__Output as _grpc_NoteId__Output } from '../grpc/NoteId';
import type { UserId as _grpc_UserId, UserId__Output as _grpc_UserId__Output } from '../grpc/UserId';

export interface NotesServiceClient extends grpc.Client {
  CreateNote(argument: _grpc_Note, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_Note__Output>): grpc.ClientUnaryCall;
  CreateNote(argument: _grpc_Note, metadata: grpc.Metadata, callback: grpc.requestCallback<_grpc_Note__Output>): grpc.ClientUnaryCall;
  CreateNote(argument: _grpc_Note, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_Note__Output>): grpc.ClientUnaryCall;
  CreateNote(argument: _grpc_Note, callback: grpc.requestCallback<_grpc_Note__Output>): grpc.ClientUnaryCall;
  createNote(argument: _grpc_Note, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_Note__Output>): grpc.ClientUnaryCall;
  createNote(argument: _grpc_Note, metadata: grpc.Metadata, callback: grpc.requestCallback<_grpc_Note__Output>): grpc.ClientUnaryCall;
  createNote(argument: _grpc_Note, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_Note__Output>): grpc.ClientUnaryCall;
  createNote(argument: _grpc_Note, callback: grpc.requestCallback<_grpc_Note__Output>): grpc.ClientUnaryCall;
  
  DeleteNote(argument: _grpc_NoteId, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_Note__Output>): grpc.ClientUnaryCall;
  DeleteNote(argument: _grpc_NoteId, metadata: grpc.Metadata, callback: grpc.requestCallback<_grpc_Note__Output>): grpc.ClientUnaryCall;
  DeleteNote(argument: _grpc_NoteId, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_Note__Output>): grpc.ClientUnaryCall;
  DeleteNote(argument: _grpc_NoteId, callback: grpc.requestCallback<_grpc_Note__Output>): grpc.ClientUnaryCall;
  deleteNote(argument: _grpc_NoteId, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_Note__Output>): grpc.ClientUnaryCall;
  deleteNote(argument: _grpc_NoteId, metadata: grpc.Metadata, callback: grpc.requestCallback<_grpc_Note__Output>): grpc.ClientUnaryCall;
  deleteNote(argument: _grpc_NoteId, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_Note__Output>): grpc.ClientUnaryCall;
  deleteNote(argument: _grpc_NoteId, callback: grpc.requestCallback<_grpc_Note__Output>): grpc.ClientUnaryCall;
  
  GetNotes(argument: _grpc_UserId, metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientReadableStream<_grpc_Note__Output>;
  GetNotes(argument: _grpc_UserId, options?: grpc.CallOptions): grpc.ClientReadableStream<_grpc_Note__Output>;
  getNotes(argument: _grpc_UserId, metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientReadableStream<_grpc_Note__Output>;
  getNotes(argument: _grpc_UserId, options?: grpc.CallOptions): grpc.ClientReadableStream<_grpc_Note__Output>;
  
}

export interface NotesServiceHandlers extends grpc.UntypedServiceImplementation {
  CreateNote: grpc.handleUnaryCall<_grpc_Note__Output, _grpc_Note>;
  
  DeleteNote: grpc.handleUnaryCall<_grpc_NoteId__Output, _grpc_Note>;
  
  GetNotes: grpc.handleServerStreamingCall<_grpc_UserId__Output, _grpc_Note>;
  
}

export interface NotesServiceDefinition extends grpc.ServiceDefinition {
  CreateNote: MethodDefinition<_grpc_Note, _grpc_Note, _grpc_Note__Output, _grpc_Note__Output>
  DeleteNote: MethodDefinition<_grpc_NoteId, _grpc_Note, _grpc_NoteId__Output, _grpc_Note__Output>
  GetNotes: MethodDefinition<_grpc_UserId, _grpc_Note, _grpc_UserId__Output, _grpc_Note__Output>
}
