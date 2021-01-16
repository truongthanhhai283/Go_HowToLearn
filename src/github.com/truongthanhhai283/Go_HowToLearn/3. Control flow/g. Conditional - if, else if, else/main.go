package main

import "fmt"

func main() {
	x := 41
	if x == 42 {
		fmt.Println(true)
	} else if x == 40 {
		fmt.Println(x)
	} else if x==39{
		fmt.Println(false)
	}else {
		fmt.Println("Cannot check condition")
	}
}
