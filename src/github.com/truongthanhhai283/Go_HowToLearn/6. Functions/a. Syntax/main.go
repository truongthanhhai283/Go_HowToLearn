package main

import "fmt"

func main() {
	x, y := demo2Returns("Hai", 28)
	fmt.Println(x)
	fmt.Println(y)

	z:=demoWithReturn("Truong Thanh Hai")
	fmt.Println(z)
}

func demo2Returns(name string, age int) (string, string) {
	a := fmt.Sprintf("Hello, my name is %s", name)
	b := fmt.Sprintf("I'm %d", age)
	return a, b
}

func demoWithReturn(name string) string {
	a:=fmt.Sprintf("Hello %s",name)
	return a
}
