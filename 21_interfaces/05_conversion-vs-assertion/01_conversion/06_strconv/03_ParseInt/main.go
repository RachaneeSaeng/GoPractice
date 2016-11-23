package main

import (
	"fmt"
	"strconv"
)

func main() {

	//	ParseBool, ParseFloat, ParseInt, and ParseUint convert strings to values:
	// if the size argument specifies a narrower width the result can be converted to that narrower type without data loss:
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-42", 10, 64) // 10 is base, 64 is bitsize
	u, _ := strconv.ParseUint("42", 10, 64)

	fmt.Println(b, f, i, u)

	//	FormatBool, FormatFloat, FormatInt, and FormatUint convert values to strings:
	w := strconv.FormatBool(true)
	x := strconv.FormatFloat(3.1415, 'E', -1, 64) // E is format, -1 is precision, 64 is bitsize
	y := strconv.FormatInt(-42, 10)               // 10 is base
	z := strconv.FormatUint(42, 16)

	fmt.Println(w, x, y, z)
}
