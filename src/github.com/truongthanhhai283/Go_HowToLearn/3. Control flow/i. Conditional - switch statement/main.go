package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println(false)

	case 2 == 2, 5 != 5 :
		fmt.Println(true)
		fallthrough

	case 3 == 3:
		fmt.Println(true)

	default:
		fmt.Println("Not really")

	}
}
