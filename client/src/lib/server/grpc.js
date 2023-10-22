import protoLoader from "@grpc/proto-loader";
import { credentials, loadPackageDefinition } from "@grpc/grpc-js";
import { ENV, SERVER_GRPC } from "$env/static/private";

export const packageDefinition = protoLoader.loadSync(
    "./src/lib/proto/main.proto",
    {
        keepCase: false,
        longs: String,
        defaults: true,
        oneofs: true,
        arrays: true,
    },
);

/** @type {import("$lib/proto/main").ProtoGrpcType} */
// @ts-ignore
// TODO: fix this
const proto = loadPackageDefinition(packageDefinition);

/** @type {import("@grpc/grpc-js").ChannelCredentials} */
const cr =
    ENV === "production"
        ? credentials.createSsl()
        : credentials.createInsecure();

export const server = new proto.proto.Service(SERVER_GRPC, cr);
