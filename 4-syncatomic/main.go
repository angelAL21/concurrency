package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("hi")

	runtime.GOMAXPROCS(4)

	var counter uint64
	var wg sync.WaitGroup

	//implementing concurrency safe counter

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("counter: %d\n", counter)
}
