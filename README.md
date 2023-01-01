# Go with SvelteKit using gRPC template

## Architecture
- SvelteKit for frontend deployed on Vercel
- Go for backend deployed on Google Cloud Run
- Backend build as microservices, with api gateway between client and server
- Communication between microservices using gRPC
- Authorization using firebase auth
- SvelteKit server for cookie and page protection, nothing is exposed to client, can work without any Javascript
- Go create firebase client at start, so no need to reconnect/recreate firebase client
- 4 calls to api with firebase authorization takes something aroung 10-20 ms :)

## Demo page
https://go-svelte-grpc.vercel.app/login

## TODO
- Emails using Pubsub
- Files using Buckets
