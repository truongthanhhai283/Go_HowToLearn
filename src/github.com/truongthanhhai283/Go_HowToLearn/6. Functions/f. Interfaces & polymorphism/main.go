package main

import (
	"fmt"
	"reflect"
)

//struct
type person struct {
	name string
	age int
}

//struct
type secretAgent struct {
	person
	DoB string
}

//method
func (p person) speak() {
	fmt.Println("I called person",p.name)
}

//method
func (s secretAgent) speak() {
	fmt.Println("I called person",s.name)
}

//func
func bar(h  human)  {
	switch h.(type) {
	case person:
		fmt.Println("I'm called human",h.(person).name)
	case secretAgent:
		fmt.Println("I'm called human",h.(secretAgent).name)
	}
}

//interface
type human interface {
	speak()
}

//create own type
type hotdog int

var a hotdog

func main() {
	s:=secretAgent{
		person: person{
			name: "Truong Thanh Hai",
			age:  25,
		},
		DoB:    "28/03/1996",
	}

	p:=person{
		name: "Nguyen Van A",
		age:  20,
	}

	s.speak()

	fmt.Println(p.name)

	//convertion
	a=100
	b:=int(a)
	fmt.Println(reflect.TypeOf(b))

	bar(s)
	bar(p)
}