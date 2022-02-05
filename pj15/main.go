package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan string, 2)

	go func(ch chan string) {
		defer close(ch)
		for i := 0; i < 5; i++ {
			msg := "message" + strconv.Itoa(i)
			ch <- msg
			fmt.Printf("SEND goroutine: %v\n", msg)
		}
	}(ch)
}
