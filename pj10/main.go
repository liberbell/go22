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
}
