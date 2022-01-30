package main

import "fmt"

func main() {
	fmt.Println("runnning main routine.")
	go g1()
	fmt.Println("exited main routine.")
}

func g1() {
	fmt.Println("running in gorouting g1.")
}
