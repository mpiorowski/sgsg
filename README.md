# Go with SvelteKit using gRPC template

## Architecture
- SvelteKit for frontend
- Go for backend
- Everything is deployed using Google Cloud Run
- Backend build as microservices, with api gateway between client and services
- Communication between microservices using gRPC
- Authorization using firebase auth
- SvelteKit server for cookie and page protection, nothing is exposed to client, can work without any Javascript
- Go create firebase client at start, so no need to reconnect/recreate firebase client

## Overview
- User can log in using google account via firebase auth
- Every call goes through api gateway, that takes care of authorizing requests
- User can add notes
- Notes sevice communicate with users services to get the user info for current note

## Demo page
https://client-dz5ydq3n2q-lz.a.run.app/

## TODO
- Emails using Pubsub
- Files using Buckets
