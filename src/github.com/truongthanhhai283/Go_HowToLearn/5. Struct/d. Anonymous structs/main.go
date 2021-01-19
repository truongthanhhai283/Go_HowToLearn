package main

import "fmt"



type person struct {
	first string
	last string
	age int
}

func main() {
	p1:=person{
		first: "Hai",
		last:  "Truong Thanh",
		age:   25,
	}

	//anonymous struct
	 p2:= struct {
		name string
		birthday string
	}{
		name:     "Hai",
		birthday: "28/3/1996",
	}

	fmt.Println(p1)

	fmt.Println(p2)
}
