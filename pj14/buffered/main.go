package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go send(ch)
	fmt.Println("\nblocking and send goroutine...")
	fmt.Printf("channel length: %v\n", len(ch))
	time.Sleep(time.Second * 2)
}

func send(ch chan string) {
	ch <- "message"
}

func receive(ch chan string) {
	time.Sleep(time.Second * 1)
}
