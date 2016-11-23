package main

import "fmt"

func main() {
	var myArray [58]int
	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", myArray)
	fmt.Printf("%T\n", mySlice)

	myArray[0] = 1000
	mySlice[0] = 1000
	fmt.Println(myArray)
	fmt.Println(mySlice)

	ChageValue(myArray[0])
	ChageValue(mySlice[0])

	fmt.Println(myArray) // myArray[0] = 1000
	fmt.Println(mySlice) // mySlice[0] = 1000

	ChageValue2(myArray)
	ChageValue3(mySlice)

	fmt.Println(myArray) // myArray[0] = 1000
	fmt.Println(mySlice) // mySlice[0] = 999
}

func ChageValue(z int) {
	z = 999
}

func ChageValue2(z [58]int) {
	z[0] = 999
}

func ChageValue3(z []int) {
	z[0] = 999
}
