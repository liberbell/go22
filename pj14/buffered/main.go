package main

import "fmt"

func main() {
	ch := make(chan string)
	go send(ch)
	fmt.Println("\nblocking and send goroutine...")
	fmt.Printf("channel length: %v\n", len(ch))
}

func send(ch chan string) {

}
