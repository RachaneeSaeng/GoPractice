package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	//	sort.StringSlice(s).Sort()
	sort.Sort(sort.StringSlice(s)) // sort slice of string
	fmt.Println(s)

}

// https://golang.org/pkg/sort/#Sort
