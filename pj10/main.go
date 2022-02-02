package main

import (
	"fmt"
	"runtime"
	"sync"
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
		fmt.Println("decrement countter = ", counter)
	}
}
