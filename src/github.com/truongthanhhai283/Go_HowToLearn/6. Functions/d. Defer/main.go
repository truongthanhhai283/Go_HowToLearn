package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func bar()  {
	fmt.Println("bar")
}
func foo()  {
	fmt.Println("foo")
}
