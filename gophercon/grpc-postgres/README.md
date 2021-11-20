# grpc-postgres

An example repo of how I like to use postgres with gRPC

```text
  +------------------------+
  |   ____  ______  ___    |
  |  /    )/      \/   \   |
  | (     / __    _\    )  |
  |  \    (/ o)  ( o)   )  |
  |   \_  (_  )   \ )  /   |
  |     \  /\_/    \)_/    |
  |      \/  //|  |\\      |
  |          v |  | v      |
  |            \__/        |
  |                        |
  +------------------------+
```

Source: [Charles Clavadetscher](https://www.postgresql.org/message-id/57386570.8090703%40swisspug.org)

## Usage

First, start a postgres container:

```bash
docker run --rm -d --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=mypass -e POSTGRES_DB=postgres postgres:13
```

Start the server:

```bash
go run main.go --postgres-url postgresql://postgres:mypass@localhost:5432/postgres
```

Navigate to [localhost](https://localhost:10000) to see the auto-generated web
UI for the service, courtesy of gRPC reflection and
[github.com/fullstorydev/grpcui](https://github.com/fullstorydev/grpcui/)!

![gRPCUI](./grpcui.png)

## Developing

### Requirements

* `go` > 1.16

### Making changes

After making any changes to the proto file or the migrations, make sure to
regenerate the files:

```bash
make generate
```

If you want to change the schema of the database, add another migration file.
Use the same naming format, with the new file names starting `2_`. Make sure to
run `make generate` and increment the migration version in `users/helpers.go`.
