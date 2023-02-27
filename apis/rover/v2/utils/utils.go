package utils

import "fmt"

func LeftPad(ref int, value, c string) string {
	ref = len(fmt.Sprint(ref))
	if ref < 2 {
		ref = 2
	}

	if ref < len(value) {
		// unlikely to ever happen
		ref = len(value)
	}

	p := padded(c, ref-len(value))
	return p + value
}

func padded(s string, size int) string {
	output := ""

	for i := 0; i < size; i++ {
		output += s
	}
	return output
}
