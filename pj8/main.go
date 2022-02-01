package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		go myfunc(i)
	}

	fmt.Println(("Each goroutine has run to competion, thanks for waiting."))
}

func myfunc(i int) {
	fmt.Println("Finished executing iteration.", i)
}
