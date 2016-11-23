package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string `json:"-"`            //tag = "-" = do not shown in marshal
	Age   int    `json:"wisdom score"` //change tag from default(Age) to "wisdom score"
}

func main() {
	p1 := person{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}
