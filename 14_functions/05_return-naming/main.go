package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe"))
	fmt.Println(ReturnId1())
	fmt.Println(ReturnId2())
}

func greet(fname string, lname string) (s string, x string) {
	s = fmt.Sprint(fname, lname)
	return
}

func ReturnId1() (id int, err error) {
	//id sent in defer function is not 10 because it pass before return
	defer func(id int) {
		if id == 10 {
			err = fmt.Errorf("Invalid Id\n")
		}
	}(id)

	id = 10

	return
}

func ReturnId2() (id int, err error) {
	//id is 10 here (use outer scope variable)
	defer func() {
		if id == 10 {
			err = fmt.Errorf("Invalid Id\n")
		}
	}()

	id = 10

	return
}

/*
IMPORTANT
Avoid using named returns.


Occasionally named returns are useful. Read this article for more information:
https://www.goinggo.net/2013/10/functions-and-naked-returns-in-go.html
*/
