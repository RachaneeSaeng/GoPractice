package main

import "fmt"

func main() {
	str := "abcge"
	fmt.Println(reverseString(str))
}

func reverseString(s string) string {
	r := []rune(s) //store ascii code in array of int32 eg. G=71
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r) //convert array of int to be string
}
