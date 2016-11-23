package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("in loop: ", i)
			c <- i // lock at i = 1 because no reader.
		}
	}()

	fmt.Println(<-c)
}

// Why does this only print zero?
// And what can you do to get it to print all 0 - 9 numbers?

// because reader read only one time
