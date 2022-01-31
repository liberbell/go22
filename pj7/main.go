package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		f := func(i int) {
			fmt.Println("i = ", i)
		}
		go f(i)
	}
	time.Sleep(10 * time.Millisecond)
}
