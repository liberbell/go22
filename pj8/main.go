package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		wg.Add(1)
		go myfunc(&wg, i)
	}
	wg.Wait()
	fmt.Println(("Each goroutine has run to competion, thanks for waiting."))
}

func myfunc(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Println("Finished executing iteration.", i)
}
