import protoLoader from '@grpc/proto-loader';
import grpc from '@grpc/grpc-js';
import type { ProtoGrpcType } from "../../grpc/grpc";
import type { UserId } from "../../grpc/grpc/UserId";
import type { Note } from "../../grpc/grpc/Note";
import { URI_USERS, URI_NOTES } from "$env/static/private";

export { UserId, Note };

export const fetchToken = async (serviceUrl: string) => {
    const tokenFetch = await fetch(
        `http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/identity?audience=https://${serviceUrl}`,
        {
            method: 'GET',
            headers: {
                'Metadata-Flavor': 'Google',
            },
        }
    );
    const token = await tokenFetch.text();
    const metadata = new grpc.Metadata();
    metadata.add('authorization', `Bearer ${token}`);
    return metadata;
};

export const packageDefinition = protoLoader.loadSync('../../grpc/grpc.proto');
export const proto = grpc.loadPackageDefinition(
    packageDefinition
) as unknown as ProtoGrpcType;

export const usersClient = new proto.grpc.UsersService(
    URI_USERS,
    grpc.credentials.createInsecure(),
);

export const notesClient = new proto.grpc.NotesService(
    URI_NOTES,
    grpc.credentials.createInsecure()
);

