
# Technical Task

The following should take about 3-4 hours.
Write a command-line utility that accepts

1. A URL
2. A destination path for a file containing a hash
3. An optional value for throttling the download
4 Given a URL
    - fetch the file, and apply throttling to the download if requested
    - without writing the file to disk hash the data using the "IMT Hash" function (described below)
    - write the hash to the specified destination file in hexadecimal format

IMT Hash description:

```go
// Length: 8 bytes
coefficients := [8]int{ 2, 3, 5, 7, 11, 13, 17, 19 }
// for each incoming byte, ib:
// for each byte of the hash, h
// h[i] = ((h[i-1] + ib) * coefficient[i]) % 255
// in the case where i-1 == -1, h[i-1] should be 0.
```

For example, hashing the data:

```go
data := []byte{12}
// Should result in a hash of:
[]byte{24, 108, 90, 204, 81, 189, 102, 126}
// When converted to hexadecimal for writing to the output file:
// 186c5acc51bd667e
```
