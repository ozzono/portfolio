# Golang microservices assignment

## Time limits

The evaluation result of the test is not linked to how much time you spend on it.
Please DO NOT spend more than ~2 hours doing it, if you haven't complete the task simply submit as is.
Successful applications show us that ~2 hours are more than enough to cover all the evaluation points below.

This assignment is meant to evaluate the golang proficiency of full-time engineers.
Your code structure should follow microservices best practices and our evaluation will focus primarily on your ability to follow good design principles and less on correctness and completeness of algorithms. During the face to face interview you will have the opportunity to explain your design choices and provide justifications for the parts that you omitted.

## Evaluation points in order of importance

- use of clean code, which is self documenting
- use of packages to achieve separation of concerns
- use of domain driven design
- use of golang idiomatic principles
- use of docker
- tests for business logic
- use of code quality checks such as linters and build tools
- use of git with appropriate commit messages
- documentation: README and inline code comments
- you MUST use go modules and a version of go >= 1.15

Results: please share a git repository with us containing your implementation.

Level of experience targeted: EXPERT

Avoid using frameworks such as go-kit and go-micro since one of the purposes of the assignment is to evaluate the candidate ability of structuring the solution in their own way.
If you have questions about the test, please draw your own conclusions.

Good luck.

Time limitations: there are no hard time limits. Once again, do not spend more than ~2 hours.

## Technical test

- Given a file with ports data (ports.json), write 2 services
- The first service (Client API) should parse the JSON file and have REST interface
- The file is of unknown size, it can contain several millions of records
The service has limited resources available (e.g. 200MB ram)
- While reading the file, it should call a second service (PortDomainService), that either creates a new record in a database, or updates the existing one
- The end result should be a database containing the ports, representing the latest version found in the JSON. Database can be Map in memory
- The first service (Client API) should provide an endpoint to retrieve the data from the second service (PortDomainService)
- Each service should be built using Dockerfile
- Provide all tests that you think are needed for your assignment. This will allow the reviewer to evaluate your critical thinking as well as your knowledge about testing
- Use gRPC as a transport between services
- The readme should explain how to run your program and test it

Choose the approach that you think is best (i.e. most flexible).

## Bonus points

- Database in docker container
- Domain Driven Design
- Docker-compose file

## Note

As mentioned earlier, the services have limited resources available, and the JSON file can be several hundred megabytes (if not gigabytes) in size.
This means that you will not able to read the entire file at once.

We are looking for the ClientAPI (the service reading the JSON) to be written in a way that is easy to reuse, give or take a few customisations.
The services themselves should handle certain signals correctly (e.g. a TERM or KILL signal should result in a graceful shutdown).
