This is an example of how to use various tools (see below) together for building Go APIs. Many of the tools and architectural choices are language-agnostic, meaning they can be used no matter the language you are using to make your API.

## Running

A dockerfile and docker-compose file are included to run the system using docker. Alternatively you could run the API manually (ie. without docker) and provide the `DSN` environment variable to point towards the correct Postgres instance. The following details how to run it locally using docker.

I use the included Taskfile in these instructions. You can [install it](https://taskfile.dev/installation/) or copy the commands within [taskfile.yaml](./taskfile.yaml). After each `task` command I include a comment listing the commands it is running.

### Reload Devstack

The simplest way to get started is the following command. It will rebuild and restart all containers, removing all data in the process. You will be prompted to apply a migration to the database.
```sh
task devstack.reload

# docker compose down -v
# docker compose up --build -d
# atlas schema apply -u $DSN --to file://schema.hcl
```
This migration is dynamically calculated by Atlas which calculates the diff between the current database state and the desired state in [schema.hcl](./schema.hcl).

### Reload Dependencies (eg. database)

If you have a need to reload just the database you can use this. You will be prompted to migrate the fresh database. See the next command for an example of why you'd need this command.
```sh
task deps.reload

# docker compose down -v
# docker compose up --build -d postgres
# sleep 2
# atlas schema apply -u $DSN --to file://schema.hcl
```

### Generate Code

When developing, you will likely need to alter the database and/or proto files. Once you've made your changes you can use the following command to generate the corresponding Go code. This will need a running, up-to-date version of the database because Jet scans it to generate Go code.
```sh
task codegen

# buf generate
# jet -dsn=$DSN -path=./gen
```
Database code is generated using [go-jet](https://github.com/go-jet/jet) (Jet), meaning it needs a running, up-to-date version of the database. Proto code is generated using [Buf](https://buf.build/) and it's dependencies. Buf is configured using [buf.gen.yaml](./buf.gen.yaml) and [buf.yaml](./buf.yaml).

## The Tools

### Connect RPC

[Connect](https://connectrpc.com/) is the foundation of the API. It's a protocol and toolkit that enables serving the API over HTTP/1.1, HTTP/2, gRPC, and gRPC-Web using a single implementation. In this project:

- API endpoints are defined in [api/proto/user/v1/user.proto](./api/proto/user/v1/user.proto)
- Generated Go code lives in [gen/user/v1/](./gen/user/v1/)
- Server implementation is in [internal/server/](./cmd/eventurely/handlers.go)

Connect RPC was chosen because:
- It provides type-safe client/server communication through Protobuf
- Supports both HTTP and gRPC protocols without additional effort
- Works seamlessly with browser-based TypeScript/JavaScript clients
- Makes integration testing straightforward since it handles HTTP marshaling

### Jet (go-jet)

[Jet](https://github.com/go-jet/jet) generates type-safe SQL builders based on your database schema. Unlike traditional ORMs, it doesn't hide the SQL - instead it provides a fluent Go API for writing queries. In this project:

- Generated models and SQL builders are in [gen/dev/public/](./gen/dev/public/)
- Repository implementations using Jet are in [internal/data/](./internal/data/)

Benefits of using Jet:
- No need to write raw SQL strings or manage manual SQL files
- Full IDE autocompletion for table names, columns, and SQL operations
- Type-safe query building that catches errors at compile time
- Clean separation between database schema and business logic
- Easy to modify queries without recompiling generated code

### Atlas

[Atlas](https://atlasgo.io/) manages our database schema and migrations using HCL (HashiCorp Configuration Language). Instead of writing SQL migrations, the desired database state is declared and Atlas handles the rest. In this project:

- Database schema is defined in [schema.hcl](./schema.hcl)
- Migration SQL is dynamically calculated
    - The SQL can be saved and checked before applying to live databases.

Key features being used:
- Declarative schema definition (WHAT state is wanted, not HOW to get there)
- Automatic migration generation based on schema changes
- Built-in validation and dry-run capabilities
- Version control friendly since schema is in a single file

### Buf

[Buf](https://buf.build/) manages the Protobuf workflow, handling code generation and schema validation. It replaces the manual protoc workflow with a more developer-friendly experience. In this project:

- Buf configuration is in [buf.yaml](./buf.yaml)
- Code generation config is in [buf.gen.yaml](./buf.gen.yaml)
- Generated code is in [gen/](./gen/)

Buf provides:
- Simplified protobuf toolchain management
- Consistent code generation across team members
- Built-in linting and breaking change detection
- Optional schema registry for sharing APIs (not used in this example; see [BSR](https://buf.build/docs/bsr/introduction/))

### Task (Taskfile)

[Task](https://taskfile.dev/) is a modern alternative to Make, providing a clean way to define and run project tasks. The [taskfile.yaml](./taskfile.yaml) includes commands for:

- `task devstack.reload`: Full environment rebuild
- `task deps.reload`: Rebuild just dependencies (e.g., database)
- `task codegen`: Generate Go code from proto files and database state

Benefits over a Makefile:
- YAML-based configuration that's easier to read, maintain, and learn
- Built-in support for dependencies between tasks
- Cross-platform compatibility without edge cases
- Clear error messages and colored output
