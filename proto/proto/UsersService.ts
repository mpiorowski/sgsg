// Original file: main.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { AuthRequest as _proto_AuthRequest, AuthRequest__Output as _proto_AuthRequest__Output } from '../proto/AuthRequest';
import type { Empty as _proto_Empty, Empty__Output as _proto_Empty__Output } from '../proto/Empty';
import type { User as _proto_User, User__Output as _proto_User__Output } from '../proto/User';
import type { UserId as _proto_UserId, UserId__Output as _proto_UserId__Output } from '../proto/UserId';

export interface UsersServiceClient extends grpc.Client {
  Auth(argument: _proto_AuthRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  Auth(argument: _proto_AuthRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  Auth(argument: _proto_AuthRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  Auth(argument: _proto_AuthRequest, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  auth(argument: _proto_AuthRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  auth(argument: _proto_AuthRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  auth(argument: _proto_AuthRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  auth(argument: _proto_AuthRequest, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  
  CreateUser(metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientDuplexStream<_proto_User, _proto_User__Output>;
  CreateUser(options?: grpc.CallOptions): grpc.ClientDuplexStream<_proto_User, _proto_User__Output>;
  createUser(metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientDuplexStream<_proto_User, _proto_User__Output>;
  createUser(options?: grpc.CallOptions): grpc.ClientDuplexStream<_proto_User, _proto_User__Output>;
  
  DeleteUser(argument: _proto_User, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  DeleteUser(argument: _proto_User, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  DeleteUser(argument: _proto_User, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  DeleteUser(argument: _proto_User, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  deleteUser(argument: _proto_User, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  deleteUser(argument: _proto_User, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  deleteUser(argument: _proto_User, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  deleteUser(argument: _proto_User, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  
  GetUser(argument: _proto_UserId, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  GetUser(argument: _proto_UserId, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  GetUser(argument: _proto_UserId, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  GetUser(argument: _proto_UserId, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  getUser(argument: _proto_UserId, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  getUser(argument: _proto_UserId, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  getUser(argument: _proto_UserId, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  getUser(argument: _proto_UserId, callback: grpc.requestCallback<_proto_User__Output>): grpc.ClientUnaryCall;
  
  GetUsers(argument: _proto_Empty, metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_User__Output>;
  GetUsers(argument: _proto_Empty, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_User__Output>;
  getUsers(argument: _proto_Empty, metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_User__Output>;
  getUsers(argument: _proto_Empty, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_User__Output>;
  
  GetUsersByIds(metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientDuplexStream<_proto_UserId, _proto_User__Output>;
  GetUsersByIds(options?: grpc.CallOptions): grpc.ClientDuplexStream<_proto_UserId, _proto_User__Output>;
  getUsersByIds(metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientDuplexStream<_proto_UserId, _proto_User__Output>;
  getUsersByIds(options?: grpc.CallOptions): grpc.ClientDuplexStream<_proto_UserId, _proto_User__Output>;
  
}

export interface UsersServiceHandlers extends grpc.UntypedServiceImplementation {
  Auth: grpc.handleUnaryCall<_proto_AuthRequest__Output, _proto_User>;
  
  CreateUser: grpc.handleBidiStreamingCall<_proto_User__Output, _proto_User>;
  
  DeleteUser: grpc.handleUnaryCall<_proto_User__Output, _proto_User>;
  
  GetUser: grpc.handleUnaryCall<_proto_UserId__Output, _proto_User>;
  
  GetUsers: grpc.handleServerStreamingCall<_proto_Empty__Output, _proto_User>;
  
  GetUsersByIds: grpc.handleBidiStreamingCall<_proto_UserId__Output, _proto_User>;
  
}

export interface UsersServiceDefinition extends grpc.ServiceDefinition {
  Auth: MethodDefinition<_proto_AuthRequest, _proto_User, _proto_AuthRequest__Output, _proto_User__Output>
  CreateUser: MethodDefinition<_proto_User, _proto_User, _proto_User__Output, _proto_User__Output>
  DeleteUser: MethodDefinition<_proto_User, _proto_User, _proto_User__Output, _proto_User__Output>
  GetUser: MethodDefinition<_proto_UserId, _proto_User, _proto_UserId__Output, _proto_User__Output>
  GetUsers: MethodDefinition<_proto_Empty, _proto_User, _proto_Empty__Output, _proto_User__Output>
  GetUsersByIds: MethodDefinition<_proto_UserId, _proto_User, _proto_UserId__Output, _proto_User__Output>
}
