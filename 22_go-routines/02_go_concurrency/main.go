package main

import (
	"fmt"
	"time"
)

// this have no waitgroup
func main() {
	go foo()
	go bar()
	time.Sleep(time.Second) // required to wait for goroutine start working unless it will finish main function imediately
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 40; i++ {
		fmt.Println("Bar:", i)
	}
}
