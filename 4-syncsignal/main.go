package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	fmt.Println("hi")
	var wg sync.WaitGroup

	mu := sync.Mutex{}
	c := sync.NewCond(&mu)
	wg.Add(1)

	go func() {
		defer wg.Done()
		//suspend goroutine until sharedRsc is populated.

		c.L.Lock()

		for len(sharedRsc) == 0 {
			c.Wait()
		}
		fmt.Println(sharedRsc["rsc1"])
		c.L.Unlock()
	}()

	c.L.Lock()
	//write changes to sharedRsc
	sharedRsc["rsc1"] = "foo"
	c.Signal()
	c.L.Unlock()
	wg.Wait()
}
