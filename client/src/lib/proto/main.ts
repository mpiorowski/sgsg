import type * as grpc from '@grpc/grpc-js';
import type { EnumTypeDefinition, MessageTypeDefinition } from '@grpc/proto-loader';

import type { AuthServiceClient as _proto_AuthServiceClient, AuthServiceDefinition as _proto_AuthServiceDefinition } from './proto/AuthService';
import type { ProfileServiceClient as _proto_ProfileServiceClient, ProfileServiceDefinition as _proto_ProfileServiceDefinition } from './proto/ProfileService';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  proto: {
    AuthResponse: MessageTypeDefinition
    AuthService: SubtypeConstructor<typeof grpc.Client, _proto_AuthServiceClient> & { service: _proto_AuthServiceDefinition }
    Count: MessageTypeDefinition
    Empty: MessageTypeDefinition
    Id: MessageTypeDefinition
    Note: MessageTypeDefinition
    Page: MessageTypeDefinition
    Profile: MessageTypeDefinition
    ProfileService: SubtypeConstructor<typeof grpc.Client, _proto_ProfileServiceClient> & { service: _proto_ProfileServiceDefinition }
    Role: EnumTypeDefinition
    StripeUrlResponse: MessageTypeDefinition
    User: MessageTypeDefinition
  }
}

