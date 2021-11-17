package main

import "fmt"

func main() {
	fmt.Println("hi")
	ch := make(chan int, 10)

	go func() { //goroutine
		defer close(ch)

		for i := 0; i < 10; i++ {
			fmt.Printf("send values to chan %d\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("received value from chan %d\n", v)
	}
}
