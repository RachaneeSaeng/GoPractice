package main

import "fmt"

func main() {
	var x [58]int
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x[42])
	x[42] = 777
	fmt.Println(x[42])
	ChageValue(x[42])
	fmt.Println(x[42])
}

func ChageValue(z int) {
	z = 999
}
