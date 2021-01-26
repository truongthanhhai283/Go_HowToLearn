package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("CPUS", runtime.NumCPU())
	fmt.Println("Goroutine", runtime.NumGoroutine())

	var wg sync.WaitGroup
	counter := 0
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutine", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutine", runtime.NumGoroutine())
	fmt.Println("count", counter)
}
