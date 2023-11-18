# SGSG - Svelte + Go + SQLite + gRPC

## What is SGSG?

It is an open-source full-stack application with two main principles in mind: **PERFORMANCE** and **SIMPLICITY**. 
The idea is that you can take this template and use it to build almost anything you need, and it will scale very well.

## Architecture
As the name suggests, the app contains four main components:
- **[SvelteKit](https://kit.svelte.dev/)** - Svelte is what I believe is currently the best frontend framework. If you've never worked with it, don't worry; it's super easy to pick up.
As an example, developers from my team who were familiar with NextJS were able to grasp it in just one week and start coding web apps. Trust me, once you try it, it's hard to go back to anything else.
- **[Go](https://go.dev/)** - The easiest backend on the market. Don't confuse simplicity with inefficiency; it's almost impossible to build a bad server using it.
- **[SQLite](https://www.sqlite.org/index.html)** - The most used database in the world. You might be skeptical about using SQLite for production, but trust me, unless you're building the next Netflix, this is all you need.
It will be faster than your PostgreSQL or MySQL because the database sits next to the backend, eliminating one network connection.
- **[gRPC](https://grpc.io/)** - Now we are delving into something a little bit harder to grasp, but I believe it's totally worth it. Two of the most important things it gives us:
    - **[Typesafety](https://protobuf.dev/)** - Thanks to protobuf, there is amazing type safety across the whole project, regardless of the language (not only for TypeScript, hi tRPC). Trust me; this is phenomenal.
  If you add one "field" to your User object, both JavaScript and Go will lint, pointing out exactly where you need to take care of it. Adding a new language like Java or Rust? Type safety for them as well.
    - **[Streaming](https://grpc.io/docs/what-is-grpc/core-concepts/#server-streaming-rpc)** - gRPC allows streaming data, which, for larger datasets, offers incredible performance. Add to this Go routines, and you can create the most amazing backend services.
  Load measurements concurrently start doing calculations on them and send them via streams as soon as any of them finish. No waiting, no blocking.

If You have any questions, feel free to ask them in Discussions or Issues. I hope this will be helpful :).

## Additional features
- **No TypeScript Build, Fully Typed with JSDocs** - Despite the absence of a TypeScript build, the code remains fully typed using JSDocs. While this approach may be somewhat controversial due to requiring more lines of code, the more I work with pure JSDocs, the more I appreciate its versatility.
It supports features like Enums, as const, and even Zod's z.infer<typeof User>, eliminating the need for the entire TypeScript build step.
- **Very Secure OAuth Implementation** - Utilizes the Phantom Token Approach with additional client-to-server authorization using an RSA key, ensuring robust security.
- **Minimal External Libraries** - Emphasizes a minimalistic approach to external libraries. From my experience, relying less on external dependencies contributes to code maintainability. This approach makes it easier to make changes even after years. It's simply the way to go.
- **Single Source of Truth Validation** - Centralizing validation on the backend simplifies logic, streamlining error checks, and ensuring a single, authoritative source for error management. Displaying these errors on the frontend remains efficient, delivering a seamless user experience.
- **Performance and Error Logging with Grafana Integration** - Efficiently log performance metrics and errors within the application, consolidating data for streamlined analysis. Utilize Grafana integration to visualize and monitor performance calls and errors, facilitating proactive management and optimization.
- **Docker for Seamless Deployment** - Leverage Docker for consistent deployment across both development and production environments. Streamline server deployment by encapsulating the application and its dependencies in containers, ensuring easy setup and scalability while maintaining environment consistency.
- **GitHub Actions for Automated Workflow** - Implement GitHub Actions to automate linting, code checks, and seamless deployments to the server. Streamline the development pipeline by integrating these actions, ensuring code quality and facilitating efficient, automatic updates to the production environment.

Thx ChatGPT for these bullet points :).

## Fast dev deployment 
In progress
