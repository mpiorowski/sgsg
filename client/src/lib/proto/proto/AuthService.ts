// Original file: proto/main.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { AuthResponse as _proto_AuthResponse, AuthResponse__Output as _proto_AuthResponse__Output } from '../proto/AuthResponse';
import type { Empty as _proto_Empty, Empty__Output as _proto_Empty__Output } from '../proto/Empty';
import type { StripeUrlResponse as _proto_StripeUrlResponse, StripeUrlResponse__Output as _proto_StripeUrlResponse__Output } from '../proto/StripeUrlResponse';

export interface AuthServiceClient extends grpc.Client {
  Auth(argument: _proto_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_AuthResponse__Output>): grpc.ClientUnaryCall;
  Auth(argument: _proto_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_AuthResponse__Output>): grpc.ClientUnaryCall;
  Auth(argument: _proto_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_AuthResponse__Output>): grpc.ClientUnaryCall;
  Auth(argument: _proto_Empty, callback: grpc.requestCallback<_proto_AuthResponse__Output>): grpc.ClientUnaryCall;
  auth(argument: _proto_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_AuthResponse__Output>): grpc.ClientUnaryCall;
  auth(argument: _proto_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_AuthResponse__Output>): grpc.ClientUnaryCall;
  auth(argument: _proto_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_AuthResponse__Output>): grpc.ClientUnaryCall;
  auth(argument: _proto_Empty, callback: grpc.requestCallback<_proto_AuthResponse__Output>): grpc.ClientUnaryCall;
  
  CreateStripeCheckout(argument: _proto_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_StripeUrlResponse__Output>): grpc.ClientUnaryCall;
  CreateStripeCheckout(argument: _proto_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_StripeUrlResponse__Output>): grpc.ClientUnaryCall;
  CreateStripeCheckout(argument: _proto_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_StripeUrlResponse__Output>): grpc.ClientUnaryCall;
  CreateStripeCheckout(argument: _proto_Empty, callback: grpc.requestCallback<_proto_StripeUrlResponse__Output>): grpc.ClientUnaryCall;
  createStripeCheckout(argument: _proto_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_StripeUrlResponse__Output>): grpc.ClientUnaryCall;
  createStripeCheckout(argument: _proto_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_StripeUrlResponse__Output>): grpc.ClientUnaryCall;
  createStripeCheckout(argument: _proto_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_StripeUrlResponse__Output>): grpc.ClientUnaryCall;
  createStripeCheckout(argument: _proto_Empty, callback: grpc.requestCallback<_proto_StripeUrlResponse__Output>): grpc.ClientUnaryCall;
  
  CreateStripePortal(argument: _proto_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_StripeUrlResponse__Output>): grpc.ClientUnaryCall;
  CreateStripePortal(argument: _proto_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_StripeUrlResponse__Output>): grpc.ClientUnaryCall;
  CreateStripePortal(argument: _proto_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_StripeUrlResponse__Output>): grpc.ClientUnaryCall;
  CreateStripePortal(argument: _proto_Empty, callback: grpc.requestCallback<_proto_StripeUrlResponse__Output>): grpc.ClientUnaryCall;
  createStripePortal(argument: _proto_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_StripeUrlResponse__Output>): grpc.ClientUnaryCall;
  createStripePortal(argument: _proto_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_StripeUrlResponse__Output>): grpc.ClientUnaryCall;
  createStripePortal(argument: _proto_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_StripeUrlResponse__Output>): grpc.ClientUnaryCall;
  createStripePortal(argument: _proto_Empty, callback: grpc.requestCallback<_proto_StripeUrlResponse__Output>): grpc.ClientUnaryCall;
  
}

export interface AuthServiceHandlers extends grpc.UntypedServiceImplementation {
  Auth: grpc.handleUnaryCall<_proto_Empty__Output, _proto_AuthResponse>;
  
  CreateStripeCheckout: grpc.handleUnaryCall<_proto_Empty__Output, _proto_StripeUrlResponse>;
  
  CreateStripePortal: grpc.handleUnaryCall<_proto_Empty__Output, _proto_StripeUrlResponse>;
  
}

export interface AuthServiceDefinition extends grpc.ServiceDefinition {
  Auth: MethodDefinition<_proto_Empty, _proto_AuthResponse, _proto_Empty__Output, _proto_AuthResponse__Output>
  CreateStripeCheckout: MethodDefinition<_proto_Empty, _proto_StripeUrlResponse, _proto_Empty__Output, _proto_StripeUrlResponse__Output>
  CreateStripePortal: MethodDefinition<_proto_Empty, _proto_StripeUrlResponse, _proto_Empty__Output, _proto_StripeUrlResponse__Output>
}
