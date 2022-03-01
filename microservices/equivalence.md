## Technical test

> I have met most of the requirements in previous codes.

1. Given a file with ports data (ports.json), write 2 services
    - [read json file example](https://github.com/ozzono/go-snippets/blob/3ccdb77860e389ad1f77eacc4a86cda67ae01434/readjson.go#L26)
    - [save into database example (with mongodb)](https://github.com/ozzono/challenges/blob/master/apis/todolist/database/database.go)
1. The first service (Client API) should parse the JSON file and have REST interface
    - [api documentation example](https://github.com/ozzono/challenges/tree/master/apis/request)
1. The file is of unknown size, it can contain several millions of records
The service has limited resources available (e.g. 200MB ram)
1. While reading the file, it should call a second service (PortDomainService), that either creates a new record in a database, or updates the existing one
    - [rest calls example](https://github.com/ozzono/challenges/blob/master/apis/request/requests/request.go)
1. The end result should be a database containing the ports, representing the latest version found in the JSON. Database can be Map in memory
1. The first service (Client API) should provide an endpoint to retrieve the data from the second service (PortDomainService)
    - [api service example](https://github.com/ozzono/challenges/blob/master/apis/request/main.go)
1. Each service should be built using Dockerfile
    - [Dockerfile sample](https://github.com/ozzono/challenges/blob/master/apis/request/Dockerfile)
    - [docker-compose.yml sample](https://github.com/ozzono/challenges/blob/master/apis/request/docker-compose.yml)
1. Provide all tests that you think are needed for your assignment. This will allow the reviewer to evaluate your critical thinking as well as your knowledge about testing
    - [project tests example](https://github.com/ozzono/challenges/tree/master/apis/request#available-tests)
1. Use gRPC as a transport between services
    - this is the only implementation I have not done by myself yet
1. The readme should explain how to run your program and test it
    - [README.md example](https://github.com/ozzono/challenges/blob/master/apis/request/README.md)

Choose the approach that you think is best (i.e. most flexible).

## Bonus points

1. Database in docker container
    - [Dockerfile sample](https://github.com/ozzono/challenges/blob/master/apis/request/Dockerfile)
1. Domain Driven Design
1. Docker-compose file
    - [docker-compose.yml sample](https://github.com/ozzono/challenges/blob/master/apis/request/docker-compose.yml)
