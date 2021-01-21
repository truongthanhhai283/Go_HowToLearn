package main

import "fmt"

func main() {
	x := map[string]int{
		"Hai": 25,
		"Bob": 30,
	}

	v, ok := x["Hai"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := x["Hai"]; ok {
		fmt.Println(true,v)
	}else {
		fmt.Println(false)
	}
}
