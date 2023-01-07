// Original file: grpc.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { AuthRequest as _grpc_AuthRequest, AuthRequest__Output as _grpc_AuthRequest__Output } from '../grpc/AuthRequest';
import type { Empty as _grpc_Empty, Empty__Output as _grpc_Empty__Output } from '../grpc/Empty';
import type { User as _grpc_User, User__Output as _grpc_User__Output } from '../grpc/User';
import type { UserId as _grpc_UserId, UserId__Output as _grpc_UserId__Output } from '../grpc/UserId';

export interface UsersServiceClient extends grpc.Client {
  Auth(argument: _grpc_AuthRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  Auth(argument: _grpc_AuthRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  Auth(argument: _grpc_AuthRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  Auth(argument: _grpc_AuthRequest, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  auth(argument: _grpc_AuthRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  auth(argument: _grpc_AuthRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  auth(argument: _grpc_AuthRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  auth(argument: _grpc_AuthRequest, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  
  CreateUser(metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientDuplexStream<_grpc_User, _grpc_User__Output>;
  CreateUser(options?: grpc.CallOptions): grpc.ClientDuplexStream<_grpc_User, _grpc_User__Output>;
  createUser(metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientDuplexStream<_grpc_User, _grpc_User__Output>;
  createUser(options?: grpc.CallOptions): grpc.ClientDuplexStream<_grpc_User, _grpc_User__Output>;
  
  DeleteUser(argument: _grpc_User, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  DeleteUser(argument: _grpc_User, metadata: grpc.Metadata, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  DeleteUser(argument: _grpc_User, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  DeleteUser(argument: _grpc_User, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  deleteUser(argument: _grpc_User, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  deleteUser(argument: _grpc_User, metadata: grpc.Metadata, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  deleteUser(argument: _grpc_User, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  deleteUser(argument: _grpc_User, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  
  GetUser(argument: _grpc_UserId, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  GetUser(argument: _grpc_UserId, metadata: grpc.Metadata, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  GetUser(argument: _grpc_UserId, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  GetUser(argument: _grpc_UserId, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  getUser(argument: _grpc_UserId, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  getUser(argument: _grpc_UserId, metadata: grpc.Metadata, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  getUser(argument: _grpc_UserId, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  getUser(argument: _grpc_UserId, callback: grpc.requestCallback<_grpc_User__Output>): grpc.ClientUnaryCall;
  
  GetUsers(argument: _grpc_Empty, metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientReadableStream<_grpc_User__Output>;
  GetUsers(argument: _grpc_Empty, options?: grpc.CallOptions): grpc.ClientReadableStream<_grpc_User__Output>;
  getUsers(argument: _grpc_Empty, metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientReadableStream<_grpc_User__Output>;
  getUsers(argument: _grpc_Empty, options?: grpc.CallOptions): grpc.ClientReadableStream<_grpc_User__Output>;
  
}

export interface UsersServiceHandlers extends grpc.UntypedServiceImplementation {
  Auth: grpc.handleUnaryCall<_grpc_AuthRequest__Output, _grpc_User>;
  
  CreateUser: grpc.handleBidiStreamingCall<_grpc_User__Output, _grpc_User>;
  
  DeleteUser: grpc.handleUnaryCall<_grpc_User__Output, _grpc_User>;
  
  GetUser: grpc.handleUnaryCall<_grpc_UserId__Output, _grpc_User>;
  
  GetUsers: grpc.handleServerStreamingCall<_grpc_Empty__Output, _grpc_User>;
  
}

export interface UsersServiceDefinition extends grpc.ServiceDefinition {
  Auth: MethodDefinition<_grpc_AuthRequest, _grpc_User, _grpc_AuthRequest__Output, _grpc_User__Output>
  CreateUser: MethodDefinition<_grpc_User, _grpc_User, _grpc_User__Output, _grpc_User__Output>
  DeleteUser: MethodDefinition<_grpc_User, _grpc_User, _grpc_User__Output, _grpc_User__Output>
  GetUser: MethodDefinition<_grpc_UserId, _grpc_User, _grpc_UserId__Output, _grpc_User__Output>
  GetUsers: MethodDefinition<_grpc_Empty, _grpc_User, _grpc_Empty__Output, _grpc_User__Output>
}
