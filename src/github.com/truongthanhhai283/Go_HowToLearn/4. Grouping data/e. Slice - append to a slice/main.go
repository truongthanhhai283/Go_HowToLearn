package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 1, 3}
	y := []int{2, 4, 1, 3, 6, 2, 5, 7, 4, 2}
	x = append(x, y...)

	fmt.Println(x)
}
