package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var counter uint64
	var wg sync.WaitGroup
	var mu sync.Mutex

	inc := func() {
		mu.Lock()
		defer mu.Unlock()
		for i := 0; i < 1000; i++ {
			counter++
		}
	}

	// dec := func() {
	// 	for i := 0; i < 1000; i++ {
	// 		counter--
	// 	}
	// }

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		func() {
			defer wg.Done()
			inc()
		}()
	}
	wg.Wait()
	fmt.Println("counter: ", counter)
}
