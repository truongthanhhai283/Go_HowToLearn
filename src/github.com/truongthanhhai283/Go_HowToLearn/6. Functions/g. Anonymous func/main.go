package main

import "fmt"

func main() {
	foo()

	//anonymous func
	func() {
		fmt.Println("Anonymous func")
	}()

	//anonymous func
	func(x int){
		fmt.Println("The meaning of life", x)
	}(32)
}

func foo() {
	fmt.Println("Food func")
}
