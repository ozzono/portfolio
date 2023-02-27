package utils

import (
	"fmt"
	"testing"
)

func TestLeftPad(t *testing.T) {

	ref := 10
	value := 1
	out := LeftPad(ref, fmt.Sprint(value), " ")
	fmt.Printf("`%s`\n", out)
	fmt.Printf("%d\n", len(out))

	ref = 1
	value = 20
	out = LeftPad(ref, fmt.Sprint(value), " ")
	fmt.Printf("`%s`\n", out)
	fmt.Printf("%d\n", len(out))
}
