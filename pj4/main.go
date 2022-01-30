package main

import "fmt"

func myFunc(step string) {
	for i := 1; 1 < 101; i++ {
		fmt.Print(step, ": ", i, " \n")
	}
}

func main() {
	myFunc
}
