package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hi")

	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "one"
	}()

	//implement timeout for rcv on channel ch
	select {
	case m := <-ch:
		fmt.Println(m)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
