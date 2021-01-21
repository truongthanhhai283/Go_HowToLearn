package main

import "fmt"

func main() {
	s := sum("x")
	fmt.Println(s)
}

//...int type: []int -> unlimited elements in slice
func sum(s string, x ...int) int {
	fmt.Println(x)
	sum := 0
	for _, v := range x {
		sum += v
	}

	fmt.Println(len(x))
	fmt.Println(cap(x))
	return sum
}
