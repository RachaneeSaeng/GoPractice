package main

import (
	"fmt"
)

func main() {
	var student []string
	var students [][]string
	//student[0] = "Todd" // this will error becuase default size is 0
	student = append(student, "Todd") //this is OK
	fmt.Println(student)
	fmt.Println(students)
}
