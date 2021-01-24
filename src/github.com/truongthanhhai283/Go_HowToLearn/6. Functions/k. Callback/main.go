package main

import "fmt"

var a int
func main() {
	s := []int{2, 1, 5, 135, 5, 12, 56, 1234, 6}
	y := sum(s...)
	fmt.Println(y)

	z := even(sum, s...)
	fmt.Println(z)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	yi := []int{}
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

