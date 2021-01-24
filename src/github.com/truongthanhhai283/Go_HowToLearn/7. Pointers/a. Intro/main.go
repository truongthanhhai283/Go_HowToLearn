package main

import "fmt"

func main() {
	var a=100
	fmt.Println(&a)
	fmt.Printf("%T\n",&a)

	b:=&a
	fmt.Println(*b)

	*b=120
	fmt.Println(a)
}
