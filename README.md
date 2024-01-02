# CQL Quickstart

In this quickstart you will see how to integrate cql within your project.

## Generate conditions

First, you will need to generate the conditions for the models using `cql-gen`.

Install `cql-gen`:

```bash
go install github.com/FrancoLiberali/cql/cql-gen@latest
```

Generate conditions:

```bash
go generate ./...
```

For details visit [cql-gen docs](github.com/FrancoLiberali/cql/cql-gen/README.md).

## Run it

First, we need a database to store the data, in this case we will use CockroachDB:

```bash
docker compose up -d
```

After that, we can run the application:

```bash
go run .
```

And you should see something like:

```bash
main.go:30
[info] Database connection is active

...

2023/08/16 10:36:28 You are ready to do queries with cql.Query[models.MyModel]
```

## Go deeper

For more details, visit [cql docs](https://compiledquerylenguage.readthedocs.io/en/latest/).
