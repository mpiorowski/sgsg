import type * as grpc from '@grpc/grpc-js';
import type { EnumTypeDefinition, MessageTypeDefinition } from '@grpc/proto-loader';

import type { FilesServiceClient as _grpc_FilesServiceClient, FilesServiceDefinition as _grpc_FilesServiceDefinition } from './grpc/FilesService';
import type { NotesServiceClient as _grpc_NotesServiceClient, NotesServiceDefinition as _grpc_NotesServiceDefinition } from './grpc/NotesService';
import type { UsersServiceClient as _grpc_UsersServiceClient, UsersServiceDefinition as _grpc_UsersServiceDefinition } from './grpc/UsersService';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  grpc: {
    AuthRequest: MessageTypeDefinition
    Empty: MessageTypeDefinition
    File: MessageTypeDefinition
    FileId: MessageTypeDefinition
    FilesService: SubtypeConstructor<typeof grpc.Client, _grpc_FilesServiceClient> & { service: _grpc_FilesServiceDefinition }
    Note: MessageTypeDefinition
    NoteId: MessageTypeDefinition
    NotesService: SubtypeConstructor<typeof grpc.Client, _grpc_NotesServiceClient> & { service: _grpc_NotesServiceDefinition }
    TargetId: MessageTypeDefinition
    User: MessageTypeDefinition
    UserId: MessageTypeDefinition
    UserRole: EnumTypeDefinition
    UsersService: SubtypeConstructor<typeof grpc.Client, _grpc_UsersServiceClient> & { service: _grpc_UsersServiceDefinition }
  }
}

