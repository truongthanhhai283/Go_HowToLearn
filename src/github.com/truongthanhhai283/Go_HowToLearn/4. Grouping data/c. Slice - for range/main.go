package main

import "fmt"

func main() {
	x := []int{1, 2, 2, 4, 5}  //slice composite literal

	for i,v:=range x{
		fmt.Println(i,v)
	}
}
