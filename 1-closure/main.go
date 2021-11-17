package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) { //passing i as int to the anonymous func
			defer wg.Done()
			fmt.Println(i)
		}(i) //passing i as argument
	}
	wg.Wait()
}
