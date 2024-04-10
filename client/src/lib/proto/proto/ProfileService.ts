// Original file: proto/main.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { Count as _proto_Count, Count__Output as _proto_Count__Output } from '../proto/Count';
import type { Empty as _proto_Empty, Empty__Output as _proto_Empty__Output } from '../proto/Empty';
import type { Id as _proto_Id, Id__Output as _proto_Id__Output } from '../proto/Id';
import type { Note as _proto_Note, Note__Output as _proto_Note__Output } from '../proto/Note';
import type { Page as _proto_Page, Page__Output as _proto_Page__Output } from '../proto/Page';
import type { Profile as _proto_Profile, Profile__Output as _proto_Profile__Output } from '../proto/Profile';

export interface ProfileServiceClient extends grpc.Client {
  CountNotesByUserId(argument: _proto_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Count__Output>): grpc.ClientUnaryCall;
  CountNotesByUserId(argument: _proto_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Count__Output>): grpc.ClientUnaryCall;
  CountNotesByUserId(argument: _proto_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Count__Output>): grpc.ClientUnaryCall;
  CountNotesByUserId(argument: _proto_Empty, callback: grpc.requestCallback<_proto_Count__Output>): grpc.ClientUnaryCall;
  countNotesByUserId(argument: _proto_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Count__Output>): grpc.ClientUnaryCall;
  countNotesByUserId(argument: _proto_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Count__Output>): grpc.ClientUnaryCall;
  countNotesByUserId(argument: _proto_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Count__Output>): grpc.ClientUnaryCall;
  countNotesByUserId(argument: _proto_Empty, callback: grpc.requestCallback<_proto_Count__Output>): grpc.ClientUnaryCall;
  
  CreateNote(argument: _proto_Note, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  CreateNote(argument: _proto_Note, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  CreateNote(argument: _proto_Note, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  CreateNote(argument: _proto_Note, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  createNote(argument: _proto_Note, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  createNote(argument: _proto_Note, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  createNote(argument: _proto_Note, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  createNote(argument: _proto_Note, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  
  DeleteNoteById(argument: _proto_Id, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  DeleteNoteById(argument: _proto_Id, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  DeleteNoteById(argument: _proto_Id, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  DeleteNoteById(argument: _proto_Id, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  deleteNoteById(argument: _proto_Id, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  deleteNoteById(argument: _proto_Id, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  deleteNoteById(argument: _proto_Id, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  deleteNoteById(argument: _proto_Id, callback: grpc.requestCallback<_proto_Empty__Output>): grpc.ClientUnaryCall;
  
  GetNoteById(argument: _proto_Id, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  GetNoteById(argument: _proto_Id, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  GetNoteById(argument: _proto_Id, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  GetNoteById(argument: _proto_Id, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  getNoteById(argument: _proto_Id, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  getNoteById(argument: _proto_Id, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  getNoteById(argument: _proto_Id, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  getNoteById(argument: _proto_Id, callback: grpc.requestCallback<_proto_Note__Output>): grpc.ClientUnaryCall;
  
  GetNotesByUserId(argument: _proto_Page, metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_Note__Output>;
  GetNotesByUserId(argument: _proto_Page, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_Note__Output>;
  getNotesByUserId(argument: _proto_Page, metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_Note__Output>;
  getNotesByUserId(argument: _proto_Page, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_Note__Output>;
  
  GetProfile(argument: _proto_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  GetProfile(argument: _proto_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  GetProfile(argument: _proto_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  GetProfile(argument: _proto_Empty, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  getProfile(argument: _proto_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  getProfile(argument: _proto_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  getProfile(argument: _proto_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  getProfile(argument: _proto_Empty, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  
  UpdateProfile(argument: _proto_Profile, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  UpdateProfile(argument: _proto_Profile, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  UpdateProfile(argument: _proto_Profile, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  UpdateProfile(argument: _proto_Profile, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  updateProfile(argument: _proto_Profile, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  updateProfile(argument: _proto_Profile, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  updateProfile(argument: _proto_Profile, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  updateProfile(argument: _proto_Profile, callback: grpc.requestCallback<_proto_Profile__Output>): grpc.ClientUnaryCall;
  
}

export interface ProfileServiceHandlers extends grpc.UntypedServiceImplementation {
  CountNotesByUserId: grpc.handleUnaryCall<_proto_Empty__Output, _proto_Count>;
  
  CreateNote: grpc.handleUnaryCall<_proto_Note__Output, _proto_Note>;
  
  DeleteNoteById: grpc.handleUnaryCall<_proto_Id__Output, _proto_Empty>;
  
  GetNoteById: grpc.handleUnaryCall<_proto_Id__Output, _proto_Note>;
  
  GetNotesByUserId: grpc.handleServerStreamingCall<_proto_Page__Output, _proto_Note>;
  
  GetProfile: grpc.handleUnaryCall<_proto_Empty__Output, _proto_Profile>;
  
  UpdateProfile: grpc.handleUnaryCall<_proto_Profile__Output, _proto_Profile>;
  
}

export interface ProfileServiceDefinition extends grpc.ServiceDefinition {
  CountNotesByUserId: MethodDefinition<_proto_Empty, _proto_Count, _proto_Empty__Output, _proto_Count__Output>
  CreateNote: MethodDefinition<_proto_Note, _proto_Note, _proto_Note__Output, _proto_Note__Output>
  DeleteNoteById: MethodDefinition<_proto_Id, _proto_Empty, _proto_Id__Output, _proto_Empty__Output>
  GetNoteById: MethodDefinition<_proto_Id, _proto_Note, _proto_Id__Output, _proto_Note__Output>
  GetNotesByUserId: MethodDefinition<_proto_Page, _proto_Note, _proto_Page__Output, _proto_Note__Output>
  GetProfile: MethodDefinition<_proto_Empty, _proto_Profile, _proto_Empty__Output, _proto_Profile__Output>
  UpdateProfile: MethodDefinition<_proto_Profile, _proto_Profile, _proto_Profile__Output, _proto_Profile__Output>
}
