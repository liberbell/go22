package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i < 4; i++ {
			time.Sleep(3 * time.Second)
			ch1 <- i
			time.Sleep(1 * time.Second)
			ch2 <- i
		}
	}()

	for i := 1; i < 5; i++ {
		select {
		case val := <-ch1:
			fmt.Printf("value received from channel ch1: %v\n", val)
		case val := <-ch2:
			fmt.Printf("value received from channel ch2: %v\n", val)
		default:
			fmt.Println("no data received... performing some other operation")
		}
		time.Sleep(2 * time.Second)
	}
}
