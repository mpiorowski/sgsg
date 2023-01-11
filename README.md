# Go with SvelteKit using gRPC template

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

## Demo page
https://client-dz5ydq3n2q-lz.a.run.app/

## Dev deployment
```
cp docker-compose.yml.dist docker-compose.yml
docker-compose up
```
