package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hi")

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 2
	}()

	//mutiplex receiv on chann ch1-ch2
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-ch1:
			fmt.Println(m1)
		case m2 := <-ch2:
			fmt.Println(m2)
		}
	}
}
