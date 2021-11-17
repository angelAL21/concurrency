package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)

	wg.Add(1)

	go func() { //goroutine 1
		defer wg.Done()

		//suspend goroutine until sharedrsc is populated

		c.L.Lock()
		for len(sharedRsc) < 1 {
			c.Wait()
		}
		fmt.Println(sharedRsc["rsc1"])
		c.L.Unlock()
	}()

	wg.Add(1)

	go func() { //goroutine 2
		defer wg.Done()
		//suspend goroutine until sharedrsc is populated
		c.L.Lock()
		for len(sharedRsc) < 2 {
			c.Wait()
		}
		fmt.Println(sharedRsc["rsc2"])
		c.L.Unlock()
	}()

	c.L.Lock()
	//Write changes to sharedrsc
	sharedRsc["rsc"] = "foo"
	sharedRsc["rsc2"] = "bar"

	//wake up goroutines
	c.Broadcast()
	c.L.Unlock()
	wg.Wait()
}
