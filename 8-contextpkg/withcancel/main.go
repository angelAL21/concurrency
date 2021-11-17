package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("hi")

	//todo: generator generates ints in a separate goroutine and sends them to the return channel.
	//callers of gen need to cancel the goroutine ounce they consume 5 ints.
	generator := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		n := 1
		go func() {
			defer close(ch)
			for {
				select {
				case ch <- n:
				case <-ctx.Done():
					return
				}
				n++
			}
		}()
		return ch
	}

	//create a context that is cancellable.
	ctx, cancel := context.WithCancel(context.Background())

	ch := generator(ctx)

	for n := range ch {
		fmt.Println(n)
		if n == 5 {
			cancel()
		}
	}
}
