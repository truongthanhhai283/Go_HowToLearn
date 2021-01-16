package main

import "fmt"

func main() {
	if true {
		fmt.Println("hello world")
	}

	if false {
		fmt.Println("hello world 1")
	}


	if !true {
		fmt.Println("hello world 2")
	}


	if !false {
		fmt.Println("hello world 3")
	}

	if x:=42;x==42 {
		fmt.Println(true)
	}
}
