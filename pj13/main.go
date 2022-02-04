package main

import "fmt"

func main() {
	ch := make(chan int)
	mult := func(x, y int) {
		res := x * y
		ch <- res
	}
	val := <-ch
	fmt.Printf("type of value of ch is: %T\n", ch)
	fmt.Printf("the value of ch is: %v\n", ch)
	fmt.Printf("the resulting value from the goroutine is: %\n", val)
}
