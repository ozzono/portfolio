# ToDo List API Code Practice

### Disclaimer

Part of the code used in this practice was inspired by the references linked below:

- [Hamza Boughraira](https://medium.com/@hamza.boughraira)'s [Hands-On With JWT in Golang](https://betterprogramming.pub/hands-on-with-jwt-in-golang-8c986d1bb4c0) tutorial;
- Golang & MongoDB usage documentation - [Quick Start: Golang & MongoDB](https://www.mongodb.com/golang):
  - [Starting and Setup](https://www.mongodb.com/blog/post/quick-start-golang--mongodb--starting-and-setup);
  - [How to Create Documents](https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-create-documents);
  - [How to Read Documents](https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-read-documents);
  - [How to Update Documents](https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-update-documents);
  - [How to Delete Documents](https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-delete-documents);
- Dealing with bson and primitives:
  - [Getting mongodb _id for go](https://dev.to/yasaricli/getting-mongodb-id-for-go-4e05);
  - [String to primitive.ObjectID](https://stackoverflow.com/questions/63879932/string-to-primitive-objectid);
  - [“golang struct to bson.d” Code Answer](https://www.codegrepper.com/code-examples/go/golang+struct+to+bson.d);
- [Docker Hub's MongoDB image usage instructions](https://hub.docker.com/_/mongo);
- [Golang: tests and working directory](https://stackoverflow.com/questions/23847003/golang-tests-and-working-directory);
- Finding MongoDB host:
  - I began to write a tool similar to [nmap](https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&cad=rja&uact=8&ved=2ahUKEwjDpYr30pbvAhUPlFkKHdPLCygQFjAAegQIARAE&url=https%3A%2F%2Fnmap.org%2F&usg=AOvVaw3VbFWhboEWZ5njba0KMQQX) before figuring out how to use docker-compose image name for internal network communication. Below, you can find the documentation that saved me and the reference used to start the address search at runtime:
  - [Networking in Compose](https://docs.docker.com/compose/networking/);
  - [Finding local addresses and masks](https://stackoverflow.com/a/53325915);
    - I kept the half-built tool for curiosity;

> All commands should be executed on repository root folder  
> Each time the project starts or the automated test are executed, the `todo` database is dropped and recreated with minimal documents.
> All generated tokens expires after 24h.

### How to use

- run `docker-compose up -d` or `make default` to start the project and release the terminal;
- run `docker-compose down` or `make stop` to stop the project;
- *to test* run `make test` or `go test ./... -v` if the services are already running;

> All authenticated endpoints accept the token sent as follows:
>
> - json payload `{"token":"data"}`
> - header `{"authorization":"data"}` or `{"authorization":"Bearer data"}`
>   - The actual format for the header configuration depends of the request origin. With Golang would be similar to this:  
>   ```req.Header.Add("authorization", "token.string.here")```

### Available Tests

- ./auth/auth_test.go
- ./database/database_test.go
- ./models/users_test.go
- ./models/requests_test.go
- ./requests/request_test.go
- ./stress/stress_test.go
- ./utils/ip_test.go

### User routes

- localhost:8888/user/add
- localhost:8888/user/get
- localhost:8888/user/update
- localhost:8888/user/login

### Request routes

- localhost:8888/request/add
- localhost:8888/request/get
- localhost:8888/request/cancel
- localhost:8888/request/approve
- localhost:8888/request/unapprove

### Unused route

> I've added the `localhost:8888/refreshdb` route for debug purposes. This route can only be used with a valid JWT.

### Known issues

The only issue I'm aware of is related to code repetition.
The MongoDB connections should be centralized to avoid it.  
As is, there are similar *CRUD* methods for `users` and `requests`.

### Test output sample

Command: `go` test -v ./...`

```log
?    todo [no test files]
=== RUN   TestNewToken
--- PASS: TestNewToken (0.00s)
=== RUN   TestValidateToken
--- PASS: TestValidateToken (0.00s)
PASS
ok   todo/auth 
=== RUN   TestDatabase
--- PASS: TestDatabase (0.01s)
PASS
ok   todo/database 
?    todo/ip [no test files]
=== RUN   TestAddDelReq
--- PASS: TestAddDelReq (0.34s)
=== RUN   TestUpdateReq
--- PASS: TestUpdateReq (0.05s)
=== RUN   TestAddDelUser
--- PASS: TestAddDelUser (0.93s)
=== RUN   TestUpdateUser
--- PASS: TestUpdateUser (2.02s)
=== RUN   TestCheckPassword
--- PASS: TestCheckPassword (2.03s)
=== RUN   TestUserHashPassword
--- PASS: TestUserHashPassword (0.89s)
PASS
ok   todo/models 
=== RUN   TestPing
--- PASS: TestPing (0.00s)
=== RUN   TestLogin
    request_test.go:48: startUser.Email   : challenge@me.more
    request_test.go:49: startUser.Password: winner lottery ticket
--- PASS: TestLogin (0.96s)
=== RUN   TestAddUser
--- PASS: TestAddUser (2.91s)
=== RUN   TestGetUser
--- PASS: TestGetUser (0.01s)
=== RUN   TestUpdateUser
--- PASS: TestUpdateUser (0.89s)
=== RUN   TestAdd
--- PASS: TestAdd (0.01s)
=== RUN   TestGet
--- PASS: TestGet (0.01s)
=== RUN   TestUpdateRequest
--- PASS: TestUpdateRequest (0.03s)
PASS
ok   todo/requests 
?    todo/route [no test files]
=== RUN   TestStress
    stress_test.go:34: Ping test
    stress_test.go:43: gimme the token
    stress_test.go:44: email   : challenge@me.more
    stress_test.go:45: password: winner lottery ticket
    stress_test.go:56: gimme 30 users
    stress_test.go:91: 30 available users
    stress_test.go:94: Ready for the first load?
    stress_test.go:123: [0000 of 1000]
    stress_test.go:123: [0100 of 1000]
    stress_test.go:117: new rps 101
    stress_test.go:123: [0200 of 1000]
    stress_test.go:117: new rps 102
    stress_test.go:117: new rps 103
    stress_test.go:123: [0300 of 1000]
    stress_test.go:117: new rps 104
    stress_test.go:117: new rps 105
    stress_test.go:123: [0400 of 1000]
    stress_test.go:123: [0500 of 1000]
    stress_test.go:123: [0600 of 1000]
    stress_test.go:117: new rps 106
    stress_test.go:123: [0700 of 1000]
    stress_test.go:123: [0800 of 1000]
    stress_test.go:123: [0900 of 1000]
    stress_test.go:117: new rps 107
    stress_test.go:127: ok, that was easy
    stress_test.go:130: 11 reqs were left behind
    stress_test.go:135: 989 available reqs
    stress_test.go:137: Ready for the second load?
    stress_test.go:159: [0000 of 0989]
    stress_test.go:159: [0100 of 0989]
    stress_test.go:155: new rps 108
    stress_test.go:159: [0200 of 0989]
    stress_test.go:159: [0300 of 0989]
    stress_test.go:159: [0400 of 0989]
    stress_test.go:155: new rps 109
    stress_test.go:159: [0500 of 0989]
    stress_test.go:159: [0600 of 0989]
    stress_test.go:155: new rps 110
    stress_test.go:155: new rps 111
    stress_test.go:159: [0700 of 0989]
    stress_test.go:159: [0800 of 0989]
    stress_test.go:155: new rps 112
    stress_test.go:159: [0900 of 0989]
    stress_test.go:172: Easy, right?
--- PASS: TestStress (34.68s)
PASS
ok   todo/stress 36.252s
=== RUN   TestFindHosts
from [127 0 0 0] to [127 255 255 255]
from [192 168 1 0] to [192 168 1 255]
from [172 17 0 0] to [172 17 255 255]
from [172 18 0 0] to [172 18 255 255]
--- PASS: TestFindHosts (0.00s)
PASS
ok   todo/utils 
```
