// Original file: ../proto/main.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { AuthResponse as _proto_AuthResponse, AuthResponse__Output as _proto_AuthResponse__Output } from '../proto/AuthResponse';
import type { Empty as _proto_Empty, Empty__Output as _proto_Empty__Output } from '../proto/Empty';
import type { Id as _proto_Id, Id__Output as _proto_Id__Output } from '../proto/Id';
import type { Note as _proto_Note, Note__Output as _proto_Note__Output } from '../proto/Note';
import type { Profile as _proto_Profile, Profile__Output as _proto_Profile__Output } from '../proto/Profile';

export interface ServiceClient extends grpc.Client {
  Auth(argument: _proto_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_AuthResponse__Output>): grpc.ClientUnaryCall;
  Auth(argument: _proto_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_AuthResponse__Output>): grpc.ClientUnaryCall;
  Auth(argument: _proto_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_AuthResponse__Output>): grpc.ClientUnaryCall;
  Auth(argument: _proto_Empty, callback: grpc.requestCallback<_proto_AuthResponse__Output>): grpc.ClientUnaryCall;
  auth(argument: _proto_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_AuthResponse__Output>): grpc.ClientUnaryCall;
  auth(argument: _proto_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_AuthResponse__Output>): grpc.ClientUnaryCall;
  auth(argument: _proto_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_AuthResponse__Output>): grpc.ClientUnaryCall;
  auth(argument: _proto_Empty, callback: grpc.requestCallback<_proto_AuthResponse__Output>): grpc.ClientUnaryCall;
  
  CreateNote(argument: _proto_Note, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  CreateNote(argument: _proto_Note, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  CreateNote(argument: _proto_Note, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  CreateNote(argument: _proto_Note, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  createNote(argument: _proto_Note, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  createNote(argument: _proto_Note, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  createNote(argument: _proto_Note, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  createNote(argument: _proto_Note, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  
  CreateProfile(argument: _proto_Profile, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  CreateProfile(argument: _proto_Profile, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  CreateProfile(argument: _proto_Profile, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  CreateProfile(argument: _proto_Profile, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  createProfile(argument: _proto_Profile, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  createProfile(argument: _proto_Profile, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  createProfile(argument: _proto_Profile, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  createProfile(argument: _proto_Profile, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  
  DeleteNoteById(argument: _proto_Id, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  DeleteNoteById(argument: _proto_Id, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  DeleteNoteById(argument: _proto_Id, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  DeleteNoteById(argument: _proto_Id, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  deleteNoteById(argument: _proto_Id, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  deleteNoteById(argument: _proto_Id, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  deleteNoteById(argument: _proto_Id, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  deleteNoteById(argument: _proto_Id, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  
  DeleteProfileById(argument: _proto_Id, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  DeleteProfileById(argument: _proto_Id, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  DeleteProfileById(argument: _proto_Id, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  DeleteProfileById(argument: _proto_Id, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  deleteProfileById(argument: _proto_Id, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  deleteProfileById(argument: _proto_Id, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  deleteProfileById(argument: _proto_Id, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  deleteProfileById(argument: _proto_Id, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  
  GetNoteById(argument: _proto_Id, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  GetNoteById(argument: _proto_Id, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  GetNoteById(argument: _proto_Id, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  GetNoteById(argument: _proto_Id, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  getNoteById(argument: _proto_Id, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  getNoteById(argument: _proto_Id, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  getNoteById(argument: _proto_Id, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  getNoteById(argument: _proto_Id, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  
  GetNotesByUserId(argument: _proto_Empty, metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_Note__Output>;
  GetNotesByUserId(argument: _proto_Empty, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_Note__Output>;
  getNotesByUserId(argument: _proto_Empty, metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_Note__Output>;
  getNotesByUserId(argument: _proto_Empty, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_Note__Output>;
  
  GetProfileByUserId(argument: _proto_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  GetProfileByUserId(argument: _proto_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  GetProfileByUserId(argument: _proto_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  GetProfileByUserId(argument: _proto_Empty, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  getProfileByUserId(argument: _proto_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  getProfileByUserId(argument: _proto_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  getProfileByUserId(argument: _proto_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  getProfileByUserId(argument: _proto_Empty, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  
}

export interface ServiceHandlers extends grpc.UntypedServiceImplementation {
  Auth: grpc.handleUnaryCall<_proto_Empty__Output, _proto_AuthResponse>;
  
  CreateNote: grpc.handleUnaryCall<_proto_Note__Output, _proto_Note>;
  
  CreateProfile: grpc.handleUnaryCall<_proto_Profile__Output, _proto_Profile>;
  
  DeleteNoteById: grpc.handleUnaryCall<_proto_Id__Output, _proto_Empty>;
  
  DeleteProfileById: grpc.handleUnaryCall<_proto_Id__Output, _proto_Empty>;
  
  GetNoteById: grpc.handleUnaryCall<_proto_Id__Output, _proto_Note>;
  
  GetNotesByUserId: grpc.handleServerStreamingCall<_proto_Empty__Output, _proto_Note>;
  
  GetProfileByUserId: grpc.handleUnaryCall<_proto_Empty__Output, _proto_Profile>;
  
}

export interface ServiceDefinition extends grpc.ServiceDefinition {
  Auth: MethodDefinition<_proto_Empty, _proto_AuthResponse, _proto_Empty__Output, _proto_AuthResponse__Output>
  CreateNote: MethodDefinition<_proto_Note, _proto_Note, _proto_Note__Output, _proto_Note__Output>
  CreateProfile: MethodDefinition<_proto_Profile, _proto_Profile, _proto_Profile__Output, _proto_Profile__Output>
  DeleteNoteById: MethodDefinition<_proto_Id, _proto_Empty, _proto_Id__Output, _proto_Empty__Output>
  DeleteProfileById: MethodDefinition<_proto_Id, _proto_Empty, _proto_Id__Output, _proto_Empty__Output>
  GetNoteById: MethodDefinition<_proto_Id, _proto_Note, _proto_Id__Output, _proto_Note__Output>
  GetNotesByUserId: MethodDefinition<_proto_Empty, _proto_Note, _proto_Empty__Output, _proto_Note__Output>
  GetProfileByUserId: MethodDefinition<_proto_Empty, _proto_Profile, _proto_Empty__Output, _proto_Profile__Output>
}
