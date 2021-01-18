package main

import "fmt"

func main() {
	x := map[string]int{
		"Hai":  30,
		"Bob":  40,
		"Jack": 50,
	}

	x["todd"] = 33

	for i, v := range x {
		fmt.Println(i, v)
	}

	y := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for k, v := range y {
		fmt.Println(k, v)
	}
}
