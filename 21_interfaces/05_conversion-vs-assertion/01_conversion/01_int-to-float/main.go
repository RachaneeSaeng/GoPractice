package main

import "fmt"

func main() {
	var x = 12
	var y = 12.1230123
	fmt.Println(y + float64(x)) // must converse before + unless it's error
	// conversion: int to float64
}
