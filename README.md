# SGSG - Svelte + Go + SQLite + gRPC

https://sgsg.bearbyte.org/

## What is SGSG?

It is an open-source full-stack application with two main principles in mind: **PERFORMANCE** and **SIMPLICITY**. 
The idea is that you can take this template and use it to build almost anything you need, and it will scale very well.

Also, this is not the next **dev** template. It has everything you need to push it to production: Nginx configuration, Docker deployments, GitHub Actions, Grafana logging, etc.

## Alternative
If you need something a little more complicated (Rust, microservices, PostgreQSL, cloud deployment), feel free to check out the second project I am running:
**[Rusve](https://github.com/mpiorowski/rusve)**

## Architecture
As the name suggests, the app contains four main components:
- **[SvelteKit](https://kit.svelte.dev/)** - Svelte currently is what I believe the best frontend framework. If you've never worked with it, don't worry; it's super easy to pick up.
As an example, developers from my team who were familiar with NextJS were able to grasp it in just one week and start coding web apps. Trust me, once you try it, it's hard to go back to anything else.
- **[Go](https://go.dev/)** - The easiest backend on the market. Don't confuse simplicity with inefficiency; it's almost impossible to build a bad server using it.
- **[SQLite](https://www.sqlite.org/index.html)** - The most used database in the world. You might be skeptical about using SQLite for production, but trust me, unless you're building the next Netflix, this is all you need.
It will be faster than your PostgreSQL or MySQL because the database sits next to the backend, eliminating one network connection. Also changing it to **PostgreSQL/MySQL** should take only a minute :).
- **[gRPC](https://grpc.io/)** - Now we are delving into something a little bit harder to grasp, but I believe it's totally worth it. Two of the most important things it gives us:
    - **[Typesafety](https://protobuf.dev/)** - Thanks to protobuf, there is amazing type safety across the whole project, regardless of the language (not only for TypeScript, hi tRPC). Trust me; this is phenomenal.
  If you add one "field" to your User object, both JavaScript and Go will lint, pointing out exactly where you need to take care of it. Adding a new language like Java or Rust? Type safety for them as well.
    - **[Streaming](https://grpc.io/docs/what-is-grpc/core-concepts/#server-streaming-rpc)** - gRPC allows streaming data, which, for larger datasets, offers incredible performance. Add to this Go routines, and you can create the most amazing backend services.
  Load measurements concurrently start doing calculations on them and send them via streams as soon as any of them finish. No waiting, no blocking.

If You have any questions, feel free to ask them in Discussions or Issues. I hope this will be helpful :).

## Additional features
- **No TypeScript Build, Fully Typed with JSDocs** - Despite the absence of a TypeScript build, the code remains fully typed using JSDocs. While this approach may be somewhat controversial due to requiring more lines of code, the more I work with pure JSDocs, the more I appreciate its versatility.
It supports features like Enums, as const, and even Zod's z.infer<typeof User>, eliminating the need for the entire TypeScript build step.
- **Fully Tested** - Backend tested using go native modules, frontend using vitest and playwright.
- **Very Secure OAuth Implementation** - Utilizes the Phantom Token Approach with additional client-to-server authorization using an RSA key, ensuring robust security.
- **Minimal External Libraries** - Emphasizes a minimalistic approach to external libraries. From my experience, relying less on external dependencies contributes to code maintainability. This approach makes it easier to make changes even after years. It's simply the way to go.
- **Single Source of Truth Validation** - Centralizing validation on the backend simplifies logic, streamlining error checks, and ensuring a single, authoritative source for error management. Displaying these errors on the frontend remains efficient, delivering a seamless user experience.
- **Performance and Error Logging with Grafana Integration** - Efficiently log performance metrics and errors within the application, consolidating data for streamlined analysis. Utilize Grafana integration to visualize and monitor performance calls and errors, facilitating proactive management and optimization.
- **Docker for Seamless Deployment** - Leverage Docker for consistent deployment across both development and production environments. Streamline server deployment by encapsulating the application and its dependencies in containers, ensuring easy setup and scalability while maintaining environment consistency.
- **GitHub Actions for Automated Workflow** - Implement GitHub Actions to automate linting, code checks, and seamless deployments to the server. Streamline the development pipeline by integrating these actions, ensuring code quality and facilitating efficient, automatic updates to the production environment.
- **Client side streaming** - Thanks to SvelteKit's newest feature, we can load and render crucial data first. Subsequently, all remaining data is returned as promises and rendered when they resolve.
- **Files, Images and Emails** - A little bit of self promotion, this application is using my another dead simple service (free) for managing files, images and emails - [UpSend](https://www.upsend.app)

Thx ChatGPT for these bullet points :).

## Test

Backend:
```
cd server
ENV=test go test ./...
```

Frontend:
```
cd client
npm run test:unit
npm run test:integration
```


## Proto

Whenever You change proto definitions, always remember to generate new types:
```
sh proto.sh
```

## Deployment

The only prerequisites are `Docker` and `Docker Compose`. 

Afterward, the only task remaining is to configure environment variables according to the deployment. No need for .env files or tedious copy/pasting — just straightforward environment variables, either configured on the system or written inline.

### Development
```
GOOGLE_CLIENT_ID=GOOGLE_CLIENT_ID \
GOOGLE_CLIENT_SECRET=GOOGLE_CLIENT_SECRET \
GITHUB_CLIENT_ID=GITHUB_CLIENT_ID \
GITHUB_CLIENT_SECRET=GITHUB_CLIENT_SECRET \
UPSEND_KEY=UPSEND_KEY \
docker compose up --build
```

### Production (or how to set up a production application for less than 10 euros per month)

Let's embark on the full journey:

1. Generate new RSA keys and push them to your repository:
```
openssl genpkey -algorithm RSA -out private.key -pkeyopt rsa_keygen_bits:2048
openssl rsa -pubout -in private.key -out public.key

mv private.key ./client/src/lib/server/private.key
mv public.key ./server/public.key
```
2. Purchase two servers (Hetzner CPX11 costs 5 euros/month each), one for the client and one for the server, with SSH key authorization.
3. Add three repository secrets to GitHub:
   - SSH_KEY: private version of the SSH key used for logging into servers
   - CLIENT_IP: client IP
   - SERVER_IP: server IP
4. Connect your domain to the servers, e.g., example.com -> client, api.example.com -> server. Also, remember to enable gRPC proxy for some providers like Cloudflare.

From here, you need to follow these instructions on both servers:

5. Log in and create a new SSH key:
```
ssh-keygen -t ed25519 -C your_email@gmail.com
```
6. Add the `.pub` version as the deployment key in your repository settings.
7. Download the repository into server.
8. Finish setting up environment variables:
    - Change domains in `docker-compose.client.yml` / `docker-compose.server.yml` 
    - Add environment variables to your system, e.g., using Fish: `set -Ux GOOGLE_CLIENT_ID 123456789`.
    - You can find all needed environment variables in `docker-compose.client.yml` / `docker-compose.server.yml`. Don't miss out on the `LOKI_URL` one — you can find it on the Grafana Cloud Portal -> Loki Details.
9. Generate certificates for your domain:
```
sudo apt-get install certbot python3-certbot-nginx
sudo certbot certonly --nginx -d example.com
sudo certbot renew --dry-run
sudo systemctl disable nginx
```
10. Install the Grafana Loki plugin for Docker:
```
docker plugin install grafana/loki-docker-driver:2.8.2 --alias loki --grant-all-permissions
```
11. Restart server.
12. Log in and run the app for the first time to ensure everything is correct:
```
cd you-repo
git checkout release
docker compose -f docker-compose.client/server.yml up -d --build
```

That's all! From now on, every push to the `release` branch will automatically pull new changes to your servers and rerun the application. You will be able to see the deployments using GitHub Actions.

Enjoy!
