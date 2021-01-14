package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println("loop ", i)
		for j := 1; j <= 3; j++ {
			fmt.Println("loop ", j)
		}
	}

	//	loop i=1++ -> loop j=1->3 ++
}
