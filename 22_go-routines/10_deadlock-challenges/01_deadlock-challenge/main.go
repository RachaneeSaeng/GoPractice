package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

// This results in a deadlock.
// Can you determine why?
// And what would you do to fix it?

// writer and reader must be in deff thread
// write will lock until reader read the channel that make this code locked in line 10
