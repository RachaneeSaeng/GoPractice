package main

import (
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person        // notice! not name embedded tpe
	LicenseToKill bool
}

func main() {
	// p1 declared in shorthand format
	p1 := doubleZero{
		person{
			"James",
			"Bond",
			20},
		true}

	// p2 declared in full format
	p2 := doubleZero{
		person: person{
			First: "Miss",
			Last:  "MoneyPenny",
			Age:   19,
		},
		LicenseToKill: false,
	} // if wanna use new line to organize braces must put comma in the prevoius line end

	fmt.Println(p1.First, p1.Last, p1.Age, p1.LicenseToKill) // don't need to call p.person.First
	fmt.Println(p2.First, p2.Last, p2.Age, p2.LicenseToKill)
}
