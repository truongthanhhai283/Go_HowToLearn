package main

import "fmt"

func main() {
	x:=make([]int,10,100)

	x[9]=999
	fmt.Println(x)

	fmt.Println(len(x))
	fmt.Println(cap(x))

	x=append(x,1000)
	fmt.Println(x)
}
