package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUS", runtime.NumCPU())
	fmt.Println("Goroutine", runtime.NumGoroutine())

	var wg sync.WaitGroup
	var counter int64
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter,1)
			fmt.Println("Counter",atomic.LoadInt64(&counter))
			v := counter
			//time.Sleep(time.Second)
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
