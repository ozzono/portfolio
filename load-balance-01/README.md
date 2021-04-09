# Load Balance Code Challenge

Challenge [description here](./customer_success_balancing/README.md)

* This challenge was totally adapted to Go from the original one in Ruby;
* I used the same data set to apply the tests;

## How to run the tests

* There are two ways of testing the challenge code;
* I easened the testing with simple `make` commands:
  * testing with go: `make go-test`
  * testing with docker: `make docker-test`

> Notes:
>
> * Testing with docker takes longer due to download and build time;
> * Neither of the test commands verifies if the needed programs are properly installed;

## Test output sample
Command: `go test -v ./...`
```log
=== RUN   TestDefault
--- PASS: TestDefault (0.00s)
=== RUN   TestCase1
--- PASS: TestCase1 (0.00s)
=== RUN   TestCase2
--- PASS: TestCase2 (0.00s)
=== RUN   TestCase3
--- PASS: TestCase3 (0.00s)
=== RUN   TestCase4
--- PASS: TestCase4 (0.00s)
=== RUN   TestCase5
--- PASS: TestCase5 (0.00s)
=== RUN   TestCase6
--- PASS: TestCase6 (0.00s)
=== RUN   TestCase7
--- PASS: TestCase7 (0.00s)
PASS
ok  	loadbalance	0.003s
```