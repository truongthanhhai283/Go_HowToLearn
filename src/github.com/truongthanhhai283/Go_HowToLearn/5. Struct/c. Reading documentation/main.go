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

type foo int

var b foo

const h = 1000

func main() {
	x, y = 4, 5

	z := secretAgent{
		person: person{
			name: "Truong Thanh Hai",
			age:  25,
		},
		ltk: false,
	}

	b=h

	fmt.Println(z)
	
	fmt.Printf("%T\n", b)
}
