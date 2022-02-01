package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var counter int
	var wg sync.WaitGroup

	inc := func() {
		counter++
		fmt.Println("increment counter = ", counter)
	}

	dec := func() {
		counter--
		fmt.Println("decrement counter = ", counter)
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				// counter++
				atomic.AddUint64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("counter: ", counter)
}
