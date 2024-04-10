import protoLoader from "@grpc/proto-loader";
import { credentials, loadPackageDefinition } from "@grpc/grpc-js";
import { env } from "$env/dynamic/private";

export const packageDefinition = protoLoader.loadSync(
    "./src/lib/proto/main.proto",
    {
        keepCase: true,
        longs: String,
        enums: Number,
        defaults: true,
        oneofs: true,
    },
);

const proto = /** @type {import("$lib/proto/main").ProtoGrpcType} */ (
    /** @type {unknown} */ (loadPackageDefinition(packageDefinition))
);

/** @type {import("@grpc/grpc-js").ChannelCredentials} */
const cr =
    env.TARGET === "production"
        ? credentials.createSsl()
        : credentials.createInsecure();

export const authService = new proto.proto.AuthService(env.AUTH_URI ?? "localhost", cr);
export const profileService = new proto.proto.ProfileService(env.PROFILE_URI ?? "localhost", cr);

/**
 * Callback function for handling gRPC responses safely.
 *
 * @template T - The type of data expected in the response.
 *
 * @param {(value: import("./safe").GrpcSafe<T>) => void} res - The callback function to handle the response.
 * @returns {(err: import("@grpc/grpc-js").ServiceError | null, data: T | undefined) => void} - A callback function to be used with gRPC response handling.
 */
export function grpcSafe(res) {
    /**
     * Handles the gRPC response and calls the provided callback function safely.
     *
     * @param {import("@grpc/grpc-js").ServiceError | null} err - The error, if any, returned in the response.
     * @param {T | undefined} data - The data returned in the response.
     */
    return (err, data) => {
        if (err) {
            if (err.code === 3) {
                let fields = [];
                try {
                    fields = JSON.parse(err.details);
                } catch (e) {
                    return res({
                        success: false,
                        error: err?.message || "Something went wrong",
                        code: err.code,
                    });
                }

                return res({
                    success: false,
                    error: "Invalid argument",
                    fields: fields,
                    code: err.code,
                });
            }
            return res({
                success: false,
                code: err.code,
                error: err?.message || "Something went wrong",
            });
        }
        if (!data) {
            return res({
                success: false,
                error: "No data returned",
                code: 0,
            });
        }
        return res({ data, success: true });
    };
}
