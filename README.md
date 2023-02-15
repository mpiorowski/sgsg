# Go with SvelteKit using gRPC template
If You have any questions, feel free to ask them in Discussions or Issues. I hope this will be helpful :).

## Architecture
![Screenshot from 2023-01-08 23-48-07](https://user-images.githubusercontent.com/26543876/211222907-97adcd78-2b81-4978-91eb-72e69c7674fc.png)

- SvelteKit for frontend and for gRPC gateway using nodejs
- Go for microservices
- SvelteKit server for page protection and data gateway, nothing is exposed to client, can work without any Javascript
- Using GCP Bucket for files and Pub/Sub for emails
- Gateway communicate with go backend using gRPC (nightmare to make it worke using Typescript...)
- Go microservices communicates with each other using also gRPC.
- So yeah, everything works on gRPC, either as streams or unary, IT IS FAST (locally request can be as fast as 3-10 ms)
- Everything is deployed using Google Cloud Run
- Authorization using Auth.js (experimental version for svelte)

Check my other similar projects:
- [Rust with SvelteKit using gRPC](https://github.com/mpiorowski/rust-grpc)
- [NodeJs with SvelteKit using GraphQL](https://github.com/mpiorowski/microservices-ts-fastify-svelte)

## Overview
- User can log in using google account via AuthJs
- Page protection on hook Handle function, so every request is protected / checked.
- Can upload/delete files stored on Bucket
- Can send email asynchronus via Pub/Sub
- Over engineered functionality of Notes:
  - Adding / deleteing notes as Unary flow, send request, get response
  - Get notes works on stream, gateway starts a stream with notes service
  - Note service get a list of notes, and for each notes, via another stream, ask users service for user data
  - Single combined note is sent to gateway as soon as it's ready
  - When all of the notes are gatehered on gateway, it return it to frontend

## Fast dev deployment (without files and email)
```
npm i --prefix client/
cp docker-compose.yml.dist docker-compose.yml
```
Fill this docker variables:
- GOOGLE_ID - google id for oauth
- GOOGLE_CLIENT - google client for oauth
- AUTH_SECRET - secret for Auth.js (https://generate-secret.vercel.app/32)
```
docker-compose up
```

## For working email and files, You need a working GCP deployment

Topic for pubsub is called "email" pointing to "/email" service POST as push PUBSUB.

Missing variables for working application:
- EMAIL_API_KEY - sendgrid api key
- EMAIL_FROM - address for email (must be verified by sendgrid)
- EMAIL_NAME - email FROM name, can be whatever You like :)
- BUCKET - name of the GCP bucket

## GCP IAM overview
![image](https://user-images.githubusercontent.com/26543876/213539599-0a4c5035-3a19-4f30-b657-a7e01ea5fcea.png)
