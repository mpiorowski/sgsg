
# Development
FROM node:20-slim AS dev

WORKDIR /client

COPY package.json /client/package.json
COPY package-lock.json /client/package-lock.json

ARG ENV
ARG PUBLIC_SERVER_HTTP
ARG SERVER_HTTP
ARG SERVER_GRPC
ARG COOKIE_DOMAIN
ARG UPSEND_KEY
RUN echo "ENV=${ENV}" >> .env
RUN echo "PUBLIC_SERVER_HTTP=${PUBLIC_SERVER_HTTP}" >> .env
RUN echo "SERVER_HTTP=${SERVER_HTTP}" >> .env
RUN echo "SERVER_GRPC=${SERVER_GRPC}" >> .env
RUN echo "COOKIE_DOMAIN=${COOKIE_DOMAIN}" >> .env
RUN echo "UPSEND_KEY=${UPSEND_KEY}" >> .env

RUN npm install

COPY . .

# Build
FROM node:20-slim AS build
WORKDIR /client

COPY package.json /client/package.json
COPY package-lock.json /client/package-lock.json

ARG ENV
ARG PUBLIC_SERVER_HTTP
ARG SERVER_HTTP
ARG SERVER_GRPC
ARG COOKIE_DOMAIN
ARG UPSEND_KEY
RUN echo "ENV=${ENV}" >> .env
RUN echo "PUBLIC_SERVER_HTTP=${PUBLIC_SERVER_HTTP}" >> .env
RUN echo "SERVER_HTTP=${SERVER_HTTP}" >> .env
RUN echo "SERVER_GRPC=${SERVER_GRPC}" >> .env
RUN echo "COOKIE_DOMAIN=${COOKIE_DOMAIN}" >> .env
RUN echo "UPSEND_KEY=${UPSEND_KEY}" >> .env

RUN npm install
COPY . .

RUN npm run build

# Production
FROM node:20-slim AS prod
WORKDIR /client

COPY package.json /client/package.json
COPY package-lock.json /client/package-lock.json
RUN npm install --production

COPY --from=build /client/build /client/build
COPY --from=build /client/src/lib/proto/ /client/src/lib/proto/
COPY --from=build /client/src/lib/server/private.key /client/src/lib/server/private.key

CMD PORT=3000 node build

