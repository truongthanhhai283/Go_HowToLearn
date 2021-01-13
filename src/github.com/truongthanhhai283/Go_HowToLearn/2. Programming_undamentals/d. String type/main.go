package main

import "fmt"

func main() {
	s := "Truong Thanh Hai"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

	for j,v:=range s{
		fmt.Printf("%d %#U",j,v);
	}
}
