package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Truong Thanh",
		Last:  "Hai",
		Age:   25,
	}

	myCrush := person{
		First: "My",
		Last:  "Crush",
		Age:   23,
	}

	people := []person{
		p1, myCrush,
	}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
