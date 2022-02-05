package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 3)
	go send(ch)
	fmt.Println("\npress enter key to continue...")
	fmt.Printf("channel length: %v\n", len(ch))
	time.Sleep(time.Second * 2)
}

func send(ch chan string) {
	ch <- "message"
}

func receive(ch chan string) {
	time.Sleep(time.Second * 1)
}
