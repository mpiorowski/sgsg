import { Metadata } from "@grpc/grpc-js";
import jwt from "jsonwebtoken";
import { env } from "$env/dynamic/private";

/**
 * Create a Metadata object with the correct authorization headers
 * Short lived token only for getting the data
 * @param {string} id - The token id
 * @returns {Metadata} - Metadata object with the correct authorization headers
 */
export function createMetadata(id) {
    const metadata = new Metadata();

    const tokenPayload = {
        id: id,
    };

    // Generate and sign the token
    const oauthToken = jwt.sign(tokenPayload, env.JWT_SECRET, {
        algorithm: "HS256",
        expiresIn: "1h",
    });

    metadata.set("x-authorization", `bearer ${oauthToken}`);
    return metadata;
}
