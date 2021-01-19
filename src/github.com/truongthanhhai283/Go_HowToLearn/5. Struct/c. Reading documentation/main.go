package main

import "fmt"

var x, y int

type person struct {
	name string
	age  int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	x, y = 4, 5

	z := secretAgent{
		person: person{
			name: "Truong Thanh Hai",
			age:  25,
		},
		ltk: false,
	}

	fmt.Println(x, y)

	fmt.Println(z)
}
