package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	f:= func(x int) {
		fmt.Println(x)
	}

	f1:= func() {
		fmt.Println("func not have params")
	}

	f(12)

	f1()
}
