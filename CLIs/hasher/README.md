# Hasher - *Code Technical Challenge*

## Table of contents

- [Hash function](#hash-function)
- [Run the code](#run-the-code)
- [Test sample](#test-sample)

## Hash function

As I did not understand what exactly was meant to be done as hash function, I implemented as near as possible of what I got figure out of the [description](./DESCRIPTION.md#L18).
> Note: by default `make` commands runs only on unix systems.

## Run the code

I compiled a binary so it can be used in any system and without requiring golang to be installed.  
Execute `./hasher -h` or `make help` for more execution instructions.

## Test sample

The test sample was executed using [this config file](/samples/config_sample-1.json).
> Note: [cmd/samples](cmd/samples) is a symbolic link to [/samples](/samples)

```log
$make test
go test ./... -v -count=1
?       hasher  [no test files]
=== RUN   TestThrottleBench
2022/05/13 23:44:40 GrabDL
2022/05/13 23:44:40 GrabDL
2022/05/13 23:44:40 unthrottled GrabDL Downloading https://apod.nasa.gov/apod/image/2205/CatsPaw_Bemmerl_960.jpg
2022/05/13 23:44:40 throttled GrabDL Downloading https://apod.nasa.gov/apod/image/2205/CatsPaw_Bemmerl_960.jpg
2022/05/13 23:44:41 - 200 OK
2022/05/13 23:44:41 - 200 OK
- throttled transferred 049152 / 150541 bytes -  32.65%
- unthrottled transferred 049152 / 150541 bytes -  32.65%
- throttled transferred 049152 / 150541 bytes -  32.65%
- unthrottled transferred 098304 / 150541 bytes -  65.30%
- throttled transferred 098304 / 150541 bytes -  65.30%
- throttled transferred 098304 / 150541 bytes -  65.30%
- throttled transferred 147456 / 150541 bytes -  97.95%
    cmd_test.go:47: unthrottled.GrabPkgDL duration -- 2443ms
    cmd_test.go:48: throttled.GrabPkgDL duration ---- 3643ms
    cmd_test.go:49: throttled request took 49.12% more time
--- PASS: TestThrottleBench (3.64s)
PASS
ok      hasher/cmd      3.651s
?       hasher/pkg/download     [no test files]
=== RUN   TestHash1
    hasher_test.go:23: hashed.Hash -- [0 101 142 18 48 112 197 47]
    hasher_test.go:24: hashed.Hex --- 00658e123070c52f
--- PASS: TestHash1 (0.00s)
=== RUN   TestHash2
    hasher_test.go:41: hashed.Hash -- [0 251 85 247 166 162 45 30]
    hasher_test.go:42: hashed.Hex --- 00fb55f7a6a22d1e
--- PASS: TestHash2 (0.00s)
PASS
ok      hasher/pkg/hasher       0.005s
=== RUN   TestFileMng
--- PASS: TestFileMng (0.00s)
PASS
ok      hasher/pkg/utils        0.008s
?       hasher/snippets/bit     [no test files]
?       hasher/snippets/copy    [no test files]
?       hasher/snippets/dustin  [no test files]
?       hasher/snippets/fatal   [no test files]
?       hasher/snippets/grab    [no test files]
?       hasher/snippets/standard        [no test files]
```
