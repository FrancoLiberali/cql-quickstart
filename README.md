# Quickstart for badaas-orm

In this quickstart you will see how to integrate badaas-orm within your project.

## Generate conditions

First, you will need to generate the conditions for the models using `badaas-cli`.

Install `badaas-cli`:

```bash
go install github.com/ditrit/badaas-cli
```

Generate conditions:

```bash
go generate ./...
```

For details visit [badaas-cli docs](github.com/ditrit/badaas-cli/README.md).

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

2023/08/16 10:36:28 You are ready to do queries with orm.NewQuery[models.MyModel]
```

## Go deeper

<!-- TODO add link to docs -->
For more details, visit [badaas-orm docs](https://github.com/ditrit/badaas/orm/README.md).
