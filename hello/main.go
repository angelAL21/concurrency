package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i <= 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	//direct call
	fun("hi, direct call")

	//goroutine function call
	go fun("goroutine1")

	//goroutine with anonymous func
	go func() {
		fun("goroutine2")
	}()

	//goroutine with func val call
	fv := fun
	go fv("goroutine3")

	//wait for goroutines to end
	time.Sleep(10 * time.Millisecond) //for goroutine1

	fmt.Println("Done... :)")
}
