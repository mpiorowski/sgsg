# Go with SvelteKit using gRPC template

## Architecture
- SvelteKit for frontend deployed on Vercel
- Go for backend deployed on Google Cloud Run
- Backend build as microservices, with api gateway between client and server
- Communication between microservices using gRPC
- Authorization using firebase auth
- SvelteKit server for cookie and route managment

## Demo page
https://go-svelte-grpc.vercel.app/login

## TODO
- Emails using Pubsub
- Files using Buckets
