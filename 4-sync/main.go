package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("hi")

	runtime.GOMAXPROCS(4)

	var balance int
	var wg sync.WaitGroup
	var mu sync.Mutex

	deposit := func(amout int) {
		mu.Lock()
		balance += amout
		mu.Unlock()
	}
	withdrawal := func(amout int) {
		mu.Lock()
		mu.Unlock()
		balance -= amout

	}
	//100 deposits of 1
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdrawal(1)
		}()
	}

	wg.Wait()
	fmt.Println(balance)

}

//mutex
