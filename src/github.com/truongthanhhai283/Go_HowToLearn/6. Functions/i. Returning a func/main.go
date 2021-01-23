package main

import "fmt"

func main() {
	fmt.Println(foo())
}

func foo() string {
	s:="Hello world"
	fmt.Println(bar()())

	return s
}

func bar() func() int  {
	return func() int {
		return 451
	}
}
