package main

import "fmt"

func doSomeThing(x int) int {
	return x * 5
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- doSomeThing(5)
	}()

	fmt.Println(<-ch)
}
