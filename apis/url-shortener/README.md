# URL shortened code challenge

## Table of contents

- [Description](#description)
- [Setting the environment](#setting-the-environment)
- [Documentation](#documentation)
- [Routes](#routes)
- [Testing](#testing)

## Description

The code follows the description as seen in the [DESCRIPTION](DESCRIPTION.md) file.

## Setting the environment

This challenge uses mongodb in a docker container, so [docker](https://docs.docker.com/engine/install/) is a requirement.  
The golang api itself is available using docker or runnning the code directly.  
Commands:

- `go run cmd/main.go`:
  - to run the code directly execute the command above;
- `make rebuild` compiles the code into a new binary file;
- `make with-docker` runs the code in docker, even with new rebuilt binary file;
- `make without-docker` runs the code outside docker and send logs to url-shortener.log;

## Documentation

I failed to make swagger work as intended, so I also wrote a PostMan Collection with the routes and added the import file in this repository.

## Routes

- ping: only pings the api and returns `pong` if it's running;
- api routes:
  - PUT: `localhost:8000/api`
    - awaits source url as json data as exampled below:

      ```json
      {
        "source":"https://go.dev"
      }
      ```

    - returns a url data as exampled below:

        ```json
        {
            "id":        "62bf0ff21d37e4cfa5608b0a",
            "source":    "https://go.dev",
            "shortened": "hkujzph",
            "count":      0,
        }
        ```

  - Get: `localhost:8000/api/:id`
    - awaits url id in path;
    - returns a url data as exampled below:

        ```json
        {
            "id":        "62bf0ff21d37e4cfa5608b0a",
            "source":    "https://go.dev",
            "shortened": "hkujzph",
            "count":      0,
        }
        ```

  - Del: `localhost:8000/api/:id`
    - awaits url id in path;
    - returns only request status code
- Redirect: `localhost:8000/<shortpath>`
  - redirects to stored shortened url; in this example:
    - shortpath: hkujzph
      - `localhost:8000/hkujzph`
    - redirects to <https://go.dev>

## Testing

> All tests and commands were executed in a linux machine.

To test it's possible to run two commands:

- `go test -v ./...`
  - if go runtime is available the above command will run all tests in verbose mode;
- `make rebuild-test` rebuild all test binary files and `make tests` run the tests binary files in verbose mode;
- test output sample:

  ```log
  $ make tests 
  ./handler.test -test.v
  === RUN   TestPing
  [GIN] 2022/07/04 - 15:27:55 | 200 |      11.314µs |                 | GET      "/ping"
  --- PASS: TestPing (0.01s)
  === RUN   TestAPIPutRoute
      handler_test.go:49: putting url
  2022/07/04 15:27:55 creating - url
  2022/07/04 15:27:55 url.ID --------- 62c3312b747c9d7aa53b609c
  2022/07/04 15:27:55 url.Source ----- https://go.dev
  2022/07/04 15:27:55 url.Shortened -- aimlwms
  2022/07/04 15:27:55 url.Count ------ 0
  [GIN] 2022/07/04 - 15:27:55 | 200 |    2.388232ms |                 | PUT      "/api"
  --- PASS: TestAPIPutRoute (0.01s)
  === RUN   TestAPIGetRoute
      handler_test.go:67: getting url
  2022/07/04 15:27:55 searching by id - url
  2022/07/04 15:27:55 url.ID --------- 62c3312b747c9d7aa53b609c
  2022/07/04 15:27:55 url.Source ----- https://go.dev
  2022/07/04 15:27:55 url.Shortened -- aimlwms
  2022/07/04 15:27:55 url.Count ------ 0
  [GIN] 2022/07/04 - 15:27:55 | 200 |     742.694µs |                 | GET      "/api/62c3312b747c9d7aa53b609c"
  --- PASS: TestAPIGetRoute (0.00s)
  === RUN   TestAPIRedirect
      handler_test.go:86: redirecting url
  2022/07/04 15:27:55 incrementing counter - url
  2022/07/04 15:27:55 url.ID --------- 62c3312b747c9d7aa53b609c
  2022/07/04 15:27:55 url.Source ----- https://go.dev
  2022/07/04 15:27:55 url.Shortened -- aimlwms
  2022/07/04 15:27:55 url.Count ------ 1
  2022/07/04 15:27:55 redirecting to https://go.dev
  [GIN] 2022/07/04 - 15:27:55 | 308 |    2.853864ms |                 | GET      "/aimlwms"
  --- PASS: TestAPIRedirect (0.03s)
  === RUN   TestAPIDelRoute
      handler_test.go:98: deleting url
  [GIN] 2022/07/04 - 15:27:55 | 200 |     586.879µs |                 | DELETE   "/api/62c3312b747c9d7aa53b609c"
  --- PASS: TestAPIDelRoute (0.00s)
  === RUN   TestAPIDelCheck
      handler_test.go:110: checking deletion
  2022/07/04 15:27:55 code: 204 msg: url not found
  [GIN] 2022/07/04 - 15:27:55 | 204 |     401.589µs |                 | GET      "/api/62c3312b747c9d7aa53b609c"
  --- PASS: TestAPIDelCheck (0.00s)
  PASS
  ./database.test -test.v
  === RUN   TestLocalHostDB
  --- PASS: TestLocalHostDB (0.00s)
  === RUN   TestURL
  2022/07/04 15:27:55 found and incremented url - url
  2022/07/04 15:27:55 url.ID --------- 62c3312bfbcb253e824e2cdb
  2022/07/04 15:27:55 url.Source ----- https://brushyo.pep
  2022/07/04 15:27:55 url.Shortened -- pwfnuk
  2022/07/04 15:27:55 url.Count ------ 1
  --- PASS: TestURL (0.01s)
  PASS
  ```
