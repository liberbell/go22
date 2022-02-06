package main

import "time"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 2
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch3 <- 3
	}()
	
	for i := 0; i < 3; i++ {
		select
	}
}
