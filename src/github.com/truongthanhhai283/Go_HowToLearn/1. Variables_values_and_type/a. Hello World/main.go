package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	foo()

	fmt.Println("something more ")
}

func foo()  {
	for i:=0;i<=100;i++{
		if i%2==0{
			fmt.Printf("i = %v\n ",i)
		}else {
			fmt.Printf("%v i/2 !=0",i)
		}
	}
}
