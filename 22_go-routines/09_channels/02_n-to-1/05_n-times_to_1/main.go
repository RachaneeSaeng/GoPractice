package main

import (
	"fmt"
)

func main() {

	n := 2
	c := make(chan int)
	done := make(chan bool)

	// n go routines to push to channel
	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
