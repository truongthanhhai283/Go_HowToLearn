package main

import "fmt"

type person struct {
	name string
	age  int
}

type secretAgent struct {
	person
	name string
	ltk bool
}

func main() {
	x := secretAgent{
		person: person{
			name: "Truong Thanh Hai",
			age:  25,
		},
		name: "Hai",
		ltk:  false,
	}

	fmt.Println(x.name)
	fmt.Println(x.person.name)
}
