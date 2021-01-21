package main

import "fmt"

func main() {
	i := 2
	for ; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("continue",i)
	}

	//for ; i <= 10; i++ {
	//	if i == 5 {
	//		break
	//	}
	//	fmt.Println("break",i)
	//}

}
