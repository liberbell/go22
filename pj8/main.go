package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		go myfunc(i)
	}
}

func myfunc(i int) {
	fmt.Println("Finished executing iteration.", i)
}
