package main

import "fmt"

func main() {
	var a [6]int
	fmt.Println("length = ",len(a))
	for _, v := range a {
		fmt.Println(v)
	}
}
