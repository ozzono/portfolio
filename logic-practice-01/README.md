# Logic Challenge

Code challenge [description here](./DESCRIPTION.md).

To run the code execute `go test -v`.

## Test output sample

Command: `go test -v`

```log
=== RUN   TestDependsAA
    logic_test.go:16: rule set is coherent
--- PASS: TestDependsAA (0.00s)
=== RUN   TestDependsAB_BA
    logic_test.go:28: rule set is coherent
--- PASS: TestDependsAB_BA (0.00s)
=== RUN   TestExclusiveAB
    logic_test.go:41: rule set is non coherent as intended: 'b' is within 'a' dependency and conflict list simultaneously
--- PASS: TestExclusiveAB (0.00s)
=== RUN   TestExclusiveAB_BC
    logic_test.go:55: rule set is non coherent as intended: 'c' is within 'a' dependency and conflict list simultaneously
--- PASS: TestExclusiveAB_BC (0.00s)
=== RUN   TestDeepDeps
    logic_test.go:72: rule set is non coherent as intended: 'f' is within 'a' dependency and conflict list simultaneously
--- PASS: TestDeepDeps (0.00s)
PASS
ok  	logic-practice-01	0.001s

```