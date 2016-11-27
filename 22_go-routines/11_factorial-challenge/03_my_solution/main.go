package main

import (
	"fmt"
	"math"
)

func main() {
	numFact := 5
	numConcurThread := 3
	var chans = make([]chan int, numConcurThread)
	chunkSize := int(math.Ceil(float64(numFact) / float64(numConcurThread)))

	for i := 0; i < numConcurThread; i++ {
		from := numFact - (i * chunkSize)
		to := numFact - ((i + 1) * chunkSize)
		if to < 1 {
			to = 1
		}
		chans[i] = factorial(from, to)
	}

	var total = 1
	for _, c := range chans {
		total *= <-c
	}
	fmt.Println(total)
}

func factorial(from, to int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := from; i > to; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
