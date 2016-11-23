package main

import "fmt"

func main() {
	var age = new([]int)   // common declaration
	bday := make([]int, 3) // shorthand declaration
	fmt.Println(age)       // new return a pointer
	fmt.Println(bday)      // make return with innitial value. Valid for slice, map and channel only
}
