package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.geekwiseacademy.com/") //http.Get return 2 objects (response and error)
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body) // assign err twice is OK
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}
