# Development
FROM node:20-slim AS dev

WORKDIR /client

COPY package.json /client/package.json
COPY package-lock.json /client/package-lock.json
RUN npm install

COPY . /client

# Build
FROM node:20-slim AS build
WORKDIR /client

COPY package.json /client/package.json
COPY package-lock.json /client/package-lock.json
RUN npm install

COPY . /client
RUN npm run build

# Production
FROM node:20-slim AS prod
WORKDIR /client

COPY package.json /client/package.json
COPY package-lock.json /client/package-lock.json
RUN npm install --omit=dev

COPY --from=build /client/build /client/build
COPY --from=build /client/src/lib/proto/ /client/src/lib/proto/

CMD PORT=3000 node build

