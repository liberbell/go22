package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("runnning main routine.")
	go g1()
	time.Sleep(10 * time.Millisecond)
	fmt.Println("exited main routine.")
}

func g1() {
	fmt.Println("running in gorouting g1.")
}
