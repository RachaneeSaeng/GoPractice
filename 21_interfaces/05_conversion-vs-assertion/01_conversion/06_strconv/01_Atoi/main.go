package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = "12"
	var y = 6
	z, _ := strconv.Atoi(x) // A to i = stirng to int
	fmt.Println(y + z)
}
