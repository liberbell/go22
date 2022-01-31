package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println("i = ", i)
		}(i)
	}
	time.Sleep(10 * time.Millisecond)
}
