# SGSG - Svelte + Go + SQLite + gRPC

## What is SGSG?

It is an open source full-stack application with one two thigns in mind, **Performance** and **Simplicity**.
The idea is that You can take this template and with it build almost anything You need. And it will scale very well.

As the name suggest, the app contains 4 main componenets:
- **[SvelteKit](https://kit.svelte.dev/)** - Svelte is what I belive right now the BEST frontend framework. Never worked with it? Don't worry, it's super easy to pick up.
Just as en example, NextJS developers in my team were able to pick it up in one week and already coding some web apps. Trust me, once You try it, it's so hard to go back to anything else.
- **[Go](https://go.dev/)** - The easiest backend on the market. And don't connect simplicty with inefficiency. It's almost impossible to build a bad server using it.
- **[SQLite](https://www.sqlite.org/index.html) - The most used database in the world. You may think, SQLite for production? Yes. Trust me. 
Unless You are bulding the next Netflix, this is all you need. And it will be faster than Your PostgreSQL or MYSql, becouse the database sit next to backend, so one less network connection.
- **[gRPC](https://grpc.io/)** - Now we are diving into one thing that is a little bit harder to catch up. But i belive it's totally worth it. Two most important things it give us:
    - **[Typesafety](https://protobuf.dev/)** - Thanks to protobuf, amazing typesafety across the whole project, no matter the language (not only for TS, hi tRPC). And trust me, this is phenomnal.
    You add one "field" to Your User object? Both JS and Go will lint, pointing excatly where You need to take care of it. Adding new languge? Java, Rust? Typesafty for them also.
    - **Stream** - gRPC allow streaming data, which for larger datasets offers increadible performance. Add to this Go routines, and You can create the most amazing backend services.
    Load measurments, concurently starts doing calculations on them and send them via streams as soon as any of them finished. No waiting, no blocking.

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
