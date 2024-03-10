import type * as grpc from '@grpc/grpc-js';
import type { EnumTypeDefinition, MessageTypeDefinition } from '@grpc/proto-loader';

import type { ServiceClient as _proto_ServiceClient, ServiceDefinition as _proto_ServiceDefinition } from './proto/Service';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  proto: {
    AuthResponse: MessageTypeDefinition
    Empty: MessageTypeDefinition
    Id: MessageTypeDefinition
    Note: MessageTypeDefinition
    Profile: MessageTypeDefinition
    Service: SubtypeConstructor<typeof grpc.Client, _proto_ServiceClient> & { service: _proto_ServiceDefinition }
    StripeUrlResponse: MessageTypeDefinition
    User: MessageTypeDefinition
    UserRole: EnumTypeDefinition
  }
}

