// Original file: grpc.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { File as _grpc_File, File__Output as _grpc_File__Output } from '../grpc/File';
import type { FileId as _grpc_FileId, FileId__Output as _grpc_FileId__Output } from '../grpc/FileId';
import type { TargetId as _grpc_TargetId, TargetId__Output as _grpc_TargetId__Output } from '../grpc/TargetId';

export interface FilesServiceClient extends grpc.Client {
  CreateFile(argument: _grpc_File, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_File__Output>): grpc.ClientUnaryCall;
  CreateFile(argument: _grpc_File, metadata: grpc.Metadata, callback: grpc.requestCallback<_grpc_File__Output>): grpc.ClientUnaryCall;
  CreateFile(argument: _grpc_File, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_File__Output>): grpc.ClientUnaryCall;
  CreateFile(argument: _grpc_File, callback: grpc.requestCallback<_grpc_File__Output>): grpc.ClientUnaryCall;
  createFile(argument: _grpc_File, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_File__Output>): grpc.ClientUnaryCall;
  createFile(argument: _grpc_File, metadata: grpc.Metadata, callback: grpc.requestCallback<_grpc_File__Output>): grpc.ClientUnaryCall;
  createFile(argument: _grpc_File, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_File__Output>): grpc.ClientUnaryCall;
  createFile(argument: _grpc_File, callback: grpc.requestCallback<_grpc_File__Output>): grpc.ClientUnaryCall;
  
  DeleteFile(argument: _grpc_FileId, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_File__Output>): grpc.ClientUnaryCall;
  DeleteFile(argument: _grpc_FileId, metadata: grpc.Metadata, callback: grpc.requestCallback<_grpc_File__Output>): grpc.ClientUnaryCall;
  DeleteFile(argument: _grpc_FileId, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_File__Output>): grpc.ClientUnaryCall;
  DeleteFile(argument: _grpc_FileId, callback: grpc.requestCallback<_grpc_File__Output>): grpc.ClientUnaryCall;
  deleteFile(argument: _grpc_FileId, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_File__Output>): grpc.ClientUnaryCall;
  deleteFile(argument: _grpc_FileId, metadata: grpc.Metadata, callback: grpc.requestCallback<_grpc_File__Output>): grpc.ClientUnaryCall;
  deleteFile(argument: _grpc_FileId, options: grpc.CallOptions, callback: grpc.requestCallback<_grpc_File__Output>): grpc.ClientUnaryCall;
  deleteFile(argument: _grpc_FileId, callback: grpc.requestCallback<_grpc_File__Output>): grpc.ClientUnaryCall;
  
  GetFiles(argument: _grpc_TargetId, metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientReadableStream<_grpc_File__Output>;
  GetFiles(argument: _grpc_TargetId, options?: grpc.CallOptions): grpc.ClientReadableStream<_grpc_File__Output>;
  getFiles(argument: _grpc_TargetId, metadata: grpc.Metadata, options?: grpc.CallOptions): grpc.ClientReadableStream<_grpc_File__Output>;
  getFiles(argument: _grpc_TargetId, options?: grpc.CallOptions): grpc.ClientReadableStream<_grpc_File__Output>;
  
}

export interface FilesServiceHandlers extends grpc.UntypedServiceImplementation {
  CreateFile: grpc.handleUnaryCall<_grpc_File__Output, _grpc_File>;
  
  DeleteFile: grpc.handleUnaryCall<_grpc_FileId__Output, _grpc_File>;
  
  GetFiles: grpc.handleServerStreamingCall<_grpc_TargetId__Output, _grpc_File>;
  
}

export interface FilesServiceDefinition extends grpc.ServiceDefinition {
  CreateFile: MethodDefinition<_grpc_File, _grpc_File, _grpc_File__Output, _grpc_File__Output>
  DeleteFile: MethodDefinition<_grpc_FileId, _grpc_File, _grpc_FileId__Output, _grpc_File__Output>
  GetFiles: MethodDefinition<_grpc_TargetId, _grpc_File, _grpc_TargetId__Output, _grpc_File__Output>
}
