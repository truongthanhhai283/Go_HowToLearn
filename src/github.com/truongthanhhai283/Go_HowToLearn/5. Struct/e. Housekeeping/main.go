package main

import "fmt"

var x int

type person struct {
	first string
	last  string
}

func main() {
	p1:=person{
		first: "Hai Truong",
		last:  "Thanh",
	}

	fmt.Println(p1)
}
