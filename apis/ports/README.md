# Ports API

Luckilly, as described [here](./DESCRIPTION.md), the proposal of this challenge is not to build and complete the challenge as a whole. Personally I believe it's not doable within 2h without a code generation tool.

## Table of contents

- [Choises](#choises)
- [What I was not able to do in time](#what-i-was-not-able-to-do-in-time)
- [Technical Test Item List](#what-i-was-not-able-to-do-in-time)
- [Running the project](#running-the-project)
- [Available Routes](#available-routes)
- [ports.json import flow](#portsjson-import-flow)
- [Note](#note)
- [Author](#author)

## Choises 
#### [back to table of contents :arrow_heading_up:](#table-of-contents)

For this challenge I made the following choises:

- PostgreSQL as database to store the ports data and make it readilly available;
- Gin gonic: I've used it before so I could develop faster with it;
- Automatic migrations with go's migration tool, making it possible to fill the database automatically on project build;
- MVC structure

## What I was not able to do in time (withing 2h) 
#### [back to table of contents :arrow_heading_up:](#table-of-contents)

- Develop the `PortDomainsService` and communicate between both services using gRPC;
- Develop automated tests for both services;

## Technical Test Item List 
#### [back to table of contents :arrow_heading_up:](#table-of-contents)
As listed in the [description](./DESCRIPTION.md)

- :green_circle: The first service (Client API) should parse the JSON file and have REST interface
  - :yellow_circle: Given a file with ports data (ports.json), write 2 services
    - The ClientAPI does read and parse the ports.json, but there is only one service finished
  - :green_circle: The file is of unknown size, it can contain several millions of records
  The service has limited resources available (e.g. 200MB ram)
  - :yellow_circle: While reading the file, it should call a second service (PortDomainService), that either creates a new record in a database, or updates the existing one
    - The ClientAPI does read and parse the ports.json, but there is only one service finished
  - :yellow_circle: The end result should be a database containing the ports, representing the latest version found in the JSON. Database can be Map in memory
    - The ClientAPI does read and parse the ports.json file:
      - Port name repetition are ignored and first registry is kept
      - Once the json is parsed, the ClientAPI keeps the last version only
  - :yellow_circle: The first service (Client API) should provide an endpoint to retrieve the data from the second service (PortDomainService)
    - As only ClientAPI was built, not data is fetched from PortDomainService
- :green_circle: Each service should be built using Dockerfile
- :red_circle: Provide all tests that you think are needed for your assignment. This will allow the reviewer to evaluate your critical thinking as well as your knowledge about testing
- :red_circle: Use gRPC as a transport between services
- :green_circle: The readme should explain how to run your program and test it

### Subtitles:
> - :green_circle: Done;
> - :red_circle: Not-done;
> - :yellow_circle: Partially done or made differently;

## Running the project 
#### [back to table of contents :arrow_heading_up:](#table-of-contents)

This project has as only requirements the docker CLI. One could clone the repo or download and unzip it.

```shell
git clone https://github.com/ozzono/challenges.git
cd challenges/api/ports
git checkout unrestricted-timeframe
make run
```

> `make run` download's the needed images; deploy the database and the golang server and add needed database strucure into postgres container.

> All new executions starts with clean database; to add data from json file visit [localhost:8000/challenge/parse-json](localhost:8000/challenge/parse-json).

Run `make stop` to interrupt the services.

## Available routes 
#### [back to table of contents :arrow_heading_up:](#table-of-contents)

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
#### [back to table of contents :arrow_heading_up:](#table-of-contents)

To import the port.json and keep it readilly available without requiring too much resources I decided to store the data in postgres database. Once the json is parsed and its data is stored, the json_controller table would have a column named as `parsed` to avoid parsing the same file multiple times.

## Note 
#### [back to table of contents :arrow_heading_up:](#table-of-contents)

- Most of the code in this project was inspired in previous experiences.
- I may have taken about 150m in total at first, including documentation.
- I worked for about 180 more minutes during the weekend.
- I worked for about 150 more minutes during the Monday and the same amount Tuesday.

## Author 
#### [back to table of contents :arrow_heading_up:](#table-of-contents)

👤 **Hugo Virgílio**

- Github [@ozzono](https://github.com/ozzono)
- LinkedIn [@hugo-virgilio](https://www.linkedin.com/in/hugo-virgilio/?locale=en_US)
