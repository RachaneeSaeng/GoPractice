package main

import (
	"fmt"
	"sync"
)

func main() {

	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2) // add statrt up number outside go func is good!

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done() //if not done waitgroup, will have error "all goroutines are asleep - deadlock!"
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
