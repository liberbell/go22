package main

import "fmt"

func main() {
	ch := make(chan string)
	go send(ch)
	fmt.Println("\nblocking and send goroutine...")
}
