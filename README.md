# SGSG - Svelte + Go + SQLite + gRPC

## What is SGSG?

It is an open-source full-stack application with two main principles in mind: **PERFORMANCE** and **SIMPLICITY**. 
The idea is that you can take this template and use it to build almost anything you need, and it will scale very well.

## Architecture
As the name suggests, the app contains four main components:
- **[SvelteKit](https://kit.svelte.dev/)** - Svelte is what I believe is currently the best frontend framework. If you've never worked with it, don't worry; it's super easy to pick up.
As an example, developers from my team who were familiar with NextJS were able to grasp it in just one week and start coding web apps. Trust me, once you try it, it's hard to go back to anything else.
- **[Go](https://go.dev/)** - The easiest backend on the market. Don't confuse simplicity with inefficiency; it's almost impossible to build a bad server using it.
- **[SQLite](https://www.sqlite.org/index.html) - The most used database in the world. You might be skeptical about using SQLite for production, but trust me, unless you're building the next Netflix, this is all you need.
It will be faster than your PostgreSQL or MySQL because the database sits next to the backend, eliminating one network connection.
- **[gRPC](https://grpc.io/)** - Now we are delving into something a little bit harder to grasp, but I believe it's totally worth it. Two of the most important things it gives us:
    - **[Typesafety](https://protobuf.dev/)** - Thanks to protobuf, there is amazing type safety across the whole project, regardless of the language (not only for TypeScript, hi tRPC). Trust me; this is phenomenal.
  If you add one "field" to your User object, both JavaScript and Go will lint, pointing out exactly where you need to take care of it. Adding a new language like Java or Rust? Type safety for them as well.
    - **Stream** - gRPC allows streaming data, which, for larger datasets, offers incredible performance. Add to this Go routines, and you can create the most amazing backend services.
  Load measurements concurrently start doing calculations on them and send them via streams as soon as any of them finish. No waiting, no blocking.

If You have any questions, feel free to ask them in Discussions or Issues. I hope this will be helpful :).

## Architecture
- SvelteKit for frontend and for gRPC gateway using nodejs
- Go for microservices
- SvelteKit server for page protection and data gateway, nothing is exposed to client, can work without any Javascript
- Using GCP Bucket for files and Pub/Sub for emails
- Gateway communicate with go backend using gRPC (nightmare to make it work using Typescript...)
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
