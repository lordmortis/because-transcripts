Because Transcripts
=====================

Because having searchable transcripts isn't just good accessibility!

# Using

# Developing

It is highly reccomended to use [docker](https://www.docker.com/) to handle installing the non-go dependencies. 

## Prerequisites

You will need a [properly setup](https://golang.org/doc/install) Go development environment to develop on this.
You will either need to use [docker](https://www.docker.com/) with the included compose script, or install a postgres 14 database.

For data model changes you'll need:

* [sqlboiler](https://github.com/volatiletech/sqlboiler#download) installed with the postgres plugin (something like `go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest`  )

## Config file setup

1. Copy `sqlboiler.sample.yaml` and `config.sample.yaml` to new files - removing the `.sample` (you should have `config.yaml` and `sqlboiler.yaml`)
2. You'll need to at least change the Database Name (`<DATABASE>`), DB User (`<DBUSER>`) and DB Password (`<DBPASS>`) in those files to a suitable config for your system.
    * It's a good idea to use a unique database and user and a random password to avoid issues with your dev environment
    * The helper tool will create a user for you
3. You may need to change the host and port entries to reflect your local setup. The example fields in `config.sample.json` are the defaults and will work with the supplied `docker-compose.yml`
4. `config.yaml` is used by the live application, while `boiler.yaml` is only used to generate new datamodels from database data.
5. The "development" flag is used for a number of things:
    * to read the migrations from files rather than the bindata
    * to read the html/email templates from files rather than the bindata
    * to output the SQL used to the console for debugging

## Bindata

You will need to build a local copy of the bindata - the helpertool will let you do this. 
See **Updating bindata via helperTool** below

## Helper tool

In the subdirectory `helperTool` of this repo is package that compiles a helper tool.
This should be built via `go build`.

- By default the helper tool will look for your config file in the working directory (the helper tool is designed to be run from `helpertool`)
- By default the helper tool will use the postgres user `postgres` (this needs to be a root user)
- By default the helper tool will use the postgres password `rootpassword`
- If you've changed the docker config these may need to be specified - run the tool with the `--help` (macOS/linux)  or `/h` (windows) option to get help

### Helper tool run examples
#### macOS / linux
```shell
cd helperTool
go build 
```

### Database Setup via helperTool

1. Setup the config file with the correct:
    1. database host
    2. database port
    3. the desired database name to use (this will be created and if this exists, it will be overwritten)
    4. the desired user to user (this will be created and if this exists, it will be overwritten)
    5. the desired user's password
2. run the helper tool with the `createDatabase` command
3. if no errors are encountered, the base setup is complete!

### Updating bindata via helperTool

1. run the helper tool with the `updateBindata` command

## Changing data models

1. Ensure you have the development flag turned on to test your migration scripts
2. run the helper tool with the `createMigration` command. This command takes the following parameters:
    1. `directory` - the root directory of this repository (if the helpertool is being run in the helperTool directory,
       then this will be either `..\` or `../` depending on your os)
    2. `name` - the name of the migration to create. Should be related to what the migration does, should probably be
       enclosed in quotes. The tool will replace whitespace with `_`, and non-alpha numeric characters with `-`
3. Specify the migration syntax in the new files created in the `datasource/migrations/` directory
4. Test your migration via compiling and running the server with the development flag turned on
5. Once your migration has the correct SQL, regenerate the bindata using the helperTool's `updateBindata` command
    1. `directory` - the root directory of this repository (if the helpertool is being run in the helperTool directory,
       then this will be either `..\` or `../` depending on your os)
6. Then generate the models:
    1. Ensure you have copied `boiler.sample.yaml` to `boiler.yaml` and changed the various directives to your local ones.
    2. Run `sqlboiler psql -o datasource_raw -p datasource_raw -c boiler.yaml` from the root of the repo.