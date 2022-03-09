# Ports API

Luckilly, as described [here](./DESCRIPTION.md), the proposal of this challenge is not to build and complete the challenge as a whole. Personally I believe it's not doable within 2h without a code generation tool.

## Table of contents

- [Choises](#choises)
- [What I was not able to do in time](#what-i-was-not-able-to-do-in-time)
- [Running the project](#running-the-project)
- [Available Routes](#available-routes)
- [ports.json import flow](#portsjson-import-flow)
- [Note](#note)
- [Author](#author)

## Choises

For this challenge I made the following choises:

- PostgreSQL as database to store the ports data and make it readilly available;
- Gin gonic: I've used it before so I could develop faster with it;
- Automatic migrations with go's migration tool, making it possible to fill the database automatically on project build;
- MVC structure

## What I was not able to do in time (withing 2h)

- Develop the `PortDomainsService` and communicate between both services using gRPC;
- Make the `ClientAPI` work properly;
- Develop automated tests for both services;
- Graceful shutdown;

## Running the project

This project has as only requirements the docker CLI. One could clone the repo or download and unzip it.

```shell
git clone https://github.com/ozzono/challenges.git
cd challenges/api/ports
git checkout unrestricted-timeframe
make run
```
> `make run` will download the needed images; deploy the database and the golang server and add needed database strucure into postgres container.

> All new executions starts with clean database; to add data from json file visit [localhost:8000/challenge/parse-json](localhost:8000/challenge/parse-json).

## Available routes

- **GET**: localhost:8000/challenge/ports/
  - fetches all ports in the database
- **GET**: localhost:8000/challenge/port/:id
  - fetches data from the port with given id
- **GET**: localhost:8000/challenge/port/code/:id
  - fetches data from the port with given code
- **POST**: localhost:8000/challenge/port
  - update data from the port with given id
- **DELETE**: localhost:8000/challenge/port/:id
  - delete port with given id
- **GET**: localhost:8000/challenge/parse-json
  - parses data from json file into the database

## ports.json import flow

To import the port.json and keep it readilly available without requiring too much resources I decided to store the data in postgres database. Once the json is parsed and its data is stored, the json_controller table would have a column named as `parsed` to avoid parsing the same file multiple times.

## Note

- Most of the code in this project was inspired in previous experiences.
- I may have taken about 150m in total at first, including documentation.
- I worked for about 180 more minutes during the weekend.
- I worked for about 150 more minutes during the Monday and the same amount Tuesday.

## Author

👤 **Hugo Virgílio**

- Github [@ozzono](https://github.com/ozzono)
- LinkedIn [@hugo-virgilio](https://www.linkedin.com/in/hugo-virgilio/?locale=en_US)
