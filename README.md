# Go with SvelteKit using gRPC template

## Architecture
- SvelteKit for frontend and for gRPC gateway using it's nodejs server
- Go for microservices
- SvelteKit server for page protection and data gateway, nothing is exposed to client, can work without any Javascript
- Svelte nodejs api gateway communicate with go backend using gRPC (nightmare to make it worke using Typescript...)
- Go microservices communicates with each other using also gRPC.
- So yeah, everything works on gRPC, either as streams or unary, IT IS FAST (locally request can be as fast as 3-10 ms)
- Everything is deployed using Google Cloud Run
- Authorization using Auth.js (experimental version for svelte)

## Overview
- User can log in using google account via AuthJs
- Page protection on hook Handle function, so every request is protected / checked.
- Over engineered functionality of Notes:
  - Adding / deleteing notes as Unary flow, send request, get response
  - Get notes works on stream, gateway starts a stream with notes service
  - Note service get a list of notes, and for each notes, via another stream, ask users service for user data
  - When all of the notes are gatehered on gateway, it return it to frontend

## Demo page
https://client-dz5ydq3n2q-lz.a.run.app/

## TODO
- Emails using Pubsub
- Files using Buckets
