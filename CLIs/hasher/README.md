## Test sample

```log
$make test
go test ./... -v -count=1
=== RUN   TestBenchMark
2022/05/12 14:08:23 GrabDL
2022/05/12 14:08:23 GrabDL
2022/05/12 14:08:23  unthrottled GrabDL Downloading https://apod.nasa.gov/apod/image/2205/CatsPaw_Bemmerl_960.jpg
2022/05/12 14:08:23  throttled GrabDL Downloading https://apod.nasa.gov/apod/image/2205/CatsPaw_Bemmerl_960.jpg
2022/05/12 14:08:25 - 200 OK
- unthrottled transferred 065536 / 150541 bytes -  43.53%
2022/05/12 14:08:26 - 200 OK
- unthrottled transferred 114688 / 150541 bytes -  76.18%
- throttled transferred 015963 / 150541 bytes -  10.60%
- throttled transferred 032768 / 150541 bytes -  21.77%
- throttled transferred 065536 / 150541 bytes -  43.53%
- throttled transferred 131072 / 150541 bytes -  87.07%
    main_test.go:56: 
    main_test.go:57: unthrottled.GrabPkgDL duration - 2770ms
    main_test.go:58:   throttled.GrabPkgDL duration - 4502ms
    main_test.go:59: throttled request took 62.53% more time
--- PASS: TestBenchMark (4.50s)
PASS
ok   hasher/cmd 4.510s
?    hasher/pkg/download [no test files]
?    hasher/pkg/utils [no test files]
```
