package main

import "fmt"

func main() {
	letter := 'A' // rune (int32) but if "A" wil be string
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)
}
