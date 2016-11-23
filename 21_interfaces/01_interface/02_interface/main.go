package main

import "fmt"

type shape interface {
	area() float64
}

type square struct {
	side float64
}

// now square type is implemented interface shape implicitly (because it have function 'area')
func (z square) area() float64 {
	return z.side * z.side
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	fmt.Printf("%T\n", s)
	info(s)
}
