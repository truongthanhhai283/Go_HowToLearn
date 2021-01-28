package main

import "fmt"

func main() {
	eye := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//	send
	go send(eye, odd, quit)

	//	receive
	receive(eye, odd, quit)

	fmt.Println("exit")
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel: ", v)

		case v := <-o:
			fmt.Println("from the even channel: ", v)

		case v := <-q:
			fmt.Println("from the even channel: ", v)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
}
