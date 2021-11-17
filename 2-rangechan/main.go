package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}
}

//unbuffered = syncronous
//buffered = given capacity, fifo queue, async.
//c:= make(chan type, capacity)
