// Squaring numbers.

package main

import (
	"fmt"
)

//generator convertes a list of int into a chan.
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

//square will square the num that we pass in the main function.
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	for n := range square(square(generator(2, 3))) {
		fmt.Println(n)
	}
}
