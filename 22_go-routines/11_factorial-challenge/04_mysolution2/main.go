package main

import (
	"fmt"
	"math"
)

func main() {
	numFact := 6
	numConcurThread := 3
	chunkSize := int(math.Ceil(float64(numFact) / float64(numConcurThread)))

	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < numConcurThread; i++ {
		from := numFact - (i * chunkSize)
		to := numFact - ((i + 1) * chunkSize)
		if to < 1 {
			to = 1
		}
		go func() {
			total := 1
			for i := from; i > to; i-- {
				total *= i
			}
			c <- total
			done <- true
		}()
	}

	go func() {
		for i := 0; i < numConcurThread; i++ {
			<-done
		}
		close(c)
	}()

	var total = 1
	for n := range c {
		total *= n
	}
	fmt.Println(total)
}
