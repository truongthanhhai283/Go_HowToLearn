package main

import "fmt"

func main() {
	s := sum(123, 123, 121, 5235, 2, 41)
	fmt.Println(s)

}

//...int type: []int -> unlimited elements in slice
func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
