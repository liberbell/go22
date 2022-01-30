package main

import (
	"fmt"
	"time"
)

func myFunc(step string) {
	for i := 1; 1 < 101; i++ {
		fmt.Print(step, ": ", i, " \n")
	}
}

func main() {
	myFunc("synchronous")
	go myFunc("G1")
	go myFunc(("Goroutine2"))
	time.Sleep(100 * time.Millisecond)
	fmt.Println("finished")
}
