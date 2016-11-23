package main

import (
	"fmt"
)

func main() {

	n := 2
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func() {
			for n := range c {
				fmt.Println(n)
			}
			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}
}
