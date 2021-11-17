package main

import "fmt"

func main() {
	fmt.Println("hi")
	ch := make(chan int)
	go func(a, b int) {
		c := a + b
		ch <- c
	}(1, 2)
	r := <-ch
	fmt.Printf("result: %d\n", r)
}
