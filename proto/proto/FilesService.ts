// Original file: main.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { File as _proto_File, File__Output as _proto_File__Output } from '../proto/File';
import type { FileId as _proto_FileId, FileId__Output as _proto_FileId__Output } from '../proto/FileId';
import type { TargetId as _proto_TargetId, TargetId__Output as _proto_TargetId__Output } from '../proto/TargetId';

export interface FilesServiceClient extends grpc.Client {
  CreateFile(argument: _proto_File, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_File__Output>): grpc.ClientUnaryCall;
  CreateFile(argument: _proto_File, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_File__Output>): grpc.ClientUnaryCall;
  CreateFile(argument: _proto_File, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_File__Output>): grpc.ClientUnaryCall;
  CreateFile(argument: _proto_File, callback: grpc.requestCallback<_proto_File__Output>): grpc.ClientUnaryCall;
  createFile(argument: _proto_File, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_File__Output>): grpc.ClientUnaryCall;
  createFile(argument: _proto_File, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_File__Output>): grpc.ClientUnaryCall;
  createFile(argument: _proto_File, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_File__Output>): grpc.ClientUnaryCall;
  createFile(argument: _proto_File, callback: grpc.requestCallback<_proto_File__Output>): grpc.ClientUnaryCall;
  
  DeleteFile(argument: _proto_FileId, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_File__Output>): grpc.ClientUnaryCall;
  DeleteFile(argument: _proto_FileId, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_File__Output>): grpc.ClientUnaryCall;
  DeleteFile(argument: _proto_FileId, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_File__Output>): grpc.ClientUnaryCall;
  DeleteFile(argument: _proto_FileId, callback: grpc.requestCallback<_proto_File__Output>): grpc.ClientUnaryCall;
  deleteFile(argument: _proto_FileId, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_File__Output>): grpc.ClientUnaryCall;
  deleteFile(argument: _proto_FileId, metadata: grpc.Metadata, callback: grpc.requestCallback<_proto_File__Output>): grpc.ClientUnaryCall;
  deleteFile(argument: _proto_FileId, options: grpc.CallOptions, callback: grpc.requestCallback<_proto_File__Output>): grpc.ClientUnaryCall;
  deleteFile(argument: _proto_FileId, callback: grpc.requestCallback<_proto_File__Output>): grpc.ClientUnaryCall;
  
  GetFiles(argument: _proto_TargetId, metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_File__Output>;
  GetFiles(argument: _proto_TargetId, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_File__Output>;
  getFiles(argument: _proto_TargetId, metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_File__Output>;
  getFiles(argument: _proto_TargetId, options?: grpc.CallOptions): grpc.ClientReadableStream<_proto_File__Output>;
  
}

export interface FilesServiceHandlers extends grpc.UntypedServiceImplementation {
  CreateFile: grpc.handleUnaryCall<_proto_File__Output, _proto_File>;
  
  DeleteFile: grpc.handleUnaryCall<_proto_FileId__Output, _proto_File>;
  
  GetFiles: grpc.handleServerStreamingCall<_proto_TargetId__Output, _proto_File>;
  
}

export interface FilesServiceDefinition extends grpc.ServiceDefinition {
  CreateFile: MethodDefinition<_proto_File, _proto_File, _proto_File__Output, _proto_File__Output>
  DeleteFile: MethodDefinition<_proto_FileId, _proto_File, _proto_FileId__Output, _proto_File__Output>
  GetFiles: MethodDefinition<_proto_TargetId, _proto_File, _proto_TargetId__Output, _proto_File__Output>
}
