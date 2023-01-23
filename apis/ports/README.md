# Ports API

Luckily, as described [here](./DESCRIPTION.md), the proposal of this challenge is not to build and complete the challenge as a whole. Without a code generation tool, it's not doable within 2h.

## Table of contents

- [Choices](#choices)
- [What I was not able to do in time](#what-i-was-not-able-to-do-in-time)
- [How to run the project](#how-to-run-the-project)
- [Running the project](#running-the-project)
- [ports.json import flow](#portsjson-import-flow)
- [Note](#note)
- [Author](#author)

## Choices

For this challenge, I made the following choices:

- PostgreSQL as a database to store the port's data and make it readily available;
- Gin gonic: I've used it before so that I could develop faster with it;
- Automatic migrations with go's migration tool, making it possible to fill the database automatically on project build;
- MVC structure

## What I was not able to do in time

- Develop the `PortDomainsService` and communicate between both services using gRPC;
- Make the `ClientAPI` work properly;
- Develop automated tests for both services;
- Graceful shutdown;

## How to run the project

Project requirements:

- docker
- docker-compose
- go >=1.16
- Run the `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest` to install the `migrate` CLI tool to run the migrations
- git CLI

### Running the project

Once the requirements are satisfied, execute the following commands to run the project:

```shell
~ git clone https://github.com/ozzono/challenges.git
~ cd challenges/apis/ports 
~ make db          # builds the postgres container
~ make migrate_up  # creates the needed tables
~ make tidy        # download package dependencies
~ make run         # runs the project locally
```

## ports.json import flow

I decided to store the data in the Postgres database to import the port.json and keep it readily available without requiring too many resources. Once the JSON is parsed and its data is stored, the json_controlle table would have a column named as `parsed` to avoid parsing the same file multiple times.

## Note

> Most of the code in this project was inspired in previous experiences.

## Author

👤 **Hugo Virgílio**

- GitHub [@ozzono](https://github.com/ozzono)
- LinkedIn [@hugo-virgilio](https://www.linkedin.com/in/hugo-virgilio/?locale=en_US)
