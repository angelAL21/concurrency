package main

import "fmt"

func main() {
	owner := func() <-chan int {
		ch := make(chan int)

		go func() {
			defer close(ch)
			for i := 0; i < 5; i++ {
				ch <- i
			}
		}()
		return ch
	}
	consumer := func(ch <-chan int) {
		//read values from chann
		for v := range ch {
			fmt.Printf("received: %v\n", v)
		}
		fmt.Println("done receiving :)")
	}
	ch := owner()
	consumer(ch)
}
