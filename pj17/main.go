package main

import "time"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i < 4; i++ {
			time.Sleep(3 * time.Second)
			ch1 <- i
			time.Sleep(1 * time.Second)
		}
	}()
}
