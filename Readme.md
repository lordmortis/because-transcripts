Because Transcripts
=====================

Because having searchable transcripts isn't just good accessibility!

# Using


# Developing

## Prerequisites

You will need a [properly setup](https://golang.org/doc/install) Go development environment to develop on this.

For data model changes you'll need:
* [sqlboiler](https://github.com/volatiletech/sqlboiler#download) installed with the SQLite3 plugin (something like `go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@latest`  )

## Bindata

You will need to build a local copy of the bindata - the helpertool will let you do this:

macOS / linux
```shell
cd helperTool
go build 
./helperTool updateBindata
```

Then the source will build properly

## Changing data models

1. Ensure you have the development flag turned on to test your migration scripts
2. run the helper tool with the `createMigration` command. This command takes the following parameters:
    1. `directory` - the root directory of this repository (if the helpertool is being run in the helperTool directory,
       then this will be either `..\` or `../` depending on your os)
    2. `name` - the name of the migration to create. Should be related to what the migration does, should probably be
       enclosed in quotes. The tool will replace whitespace with `_`, and non-alpha numeric characters with `-`
3. Specify the migration syntax in the new files created in the `datasource/migrations/` directory
4. Test your migration via either:
    1. compiling and running the server with the development flag turned on
    2. manually running the migration via `migrate -source file://datasource/migrations -database postgres://<DBUSER>:<DBPASS>@localhost:15432/<DBNAME>?sslmode=disable up`
5. Once your migration has the correct SQL, regenerate the bindata using the helperTool's `updateBindata` command
    1. `directory` - the root directory of this repository (if the helpertool is being run in the helperTool directory,
       then this will be either `..\` or `../` depending on your os)
6. Then generate the models:
    1. Ensure you have copied `boiler.sample.yaml` to `boiler.yaml` and changed the various directives to your local ones.
    2. Run `sqlboiler sqlite3 -o datasources_raw -p datasource_raw -c boiler.yaml` from the root of the repo.