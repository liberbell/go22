package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	ch := make(chan string, 3)
	go send(ch)
	fmt.Println("\npress enter key to continue...")
	fmt.Scanln()
	go receive(ch)
	time.Sleep(time.Second * 1)
}

func send(ch chan string) {
	fmt.Printf("channel length before sending anything: %v\n", len(ch))
	for i := 0; i < 4; i++ {
		msg := "message" + strconv.Itoa(i)
	}
}

func receive(ch chan string) {
	time.Sleep(time.Second * 1)
}
