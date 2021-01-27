package main

import "fmt"

func main() {
	c:=make(chan int)

	c<-43

	go func() {
		c<-32
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
}
