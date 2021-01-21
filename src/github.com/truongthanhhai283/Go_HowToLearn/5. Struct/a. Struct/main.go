package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	p1:=person{
		name:"Truong Thanh Hai",
		age: 25,
	}

	p2:=person{
		name:"Tran Van C",
		age: 20,
	}

	fmt.Println(p1)

	fmt.Println(p2.age)
}
