import protoLoader from '@grpc/proto-loader';
import grpc from '@grpc/grpc-js';
import type { ProtoGrpcType } from "../../proto/main";
import type { UserId } from "../../proto/proto/UserId";
import type { Note } from "../../proto/proto/Note";
import { URI_USERS, URI_NOTES, NODE_ENV, URI_FILES } from "$env/static/private";

export { UserId, Note };

const cacheToken = new Map<string, {
    expires: Date;
    metadata: grpc.Metadata;
}>();
export const fetchToken = async (serviceUrl: string) => {
    if (NODE_ENV === "development") {
        return new grpc.Metadata();
    }

    // check cache for token
    const cached = cacheToken.get(serviceUrl);
    if (cached && cached.expires > new Date()) {
        console.info("Using cached token");
        return cached.metadata;
    }

    console.info("Fetching token");
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

    // cache token for 1 hour
    cacheToken.set(serviceUrl, {
        expires: new Date(Date.now() + 3600000),
        metadata,
    });

    return metadata;
};

export const packageDefinition = protoLoader.loadSync('../proto/main.proto');
export const proto = grpc.loadPackageDefinition(
    packageDefinition
) as unknown as ProtoGrpcType;

export const usersClient = new proto.proto.UsersService(
    URI_USERS,
    NODE_ENV === 'production' ? grpc.credentials.createSsl() : grpc.credentials.createInsecure()
);

export const notesClient = new proto.proto.NotesService(
    URI_NOTES,
    NODE_ENV === 'production' ? grpc.credentials.createSsl() : grpc.credentials.createInsecure()
);

export const filesClient = new proto.proto.FilesService(
    URI_FILES,
    NODE_ENV === 'production' ? grpc.credentials.createSsl() : grpc.credentials.createInsecure()
);

