package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("hi")

	var data int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		data++
	}()
	wg.Wait()

	fmt.Printf("the value of data is %d\n", data)
	fmt.Println("done :)")
}
