package main

import (
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var counter uint64
	var wg sync.WaitGroup

	inc := func() {
		for i := 0; i < 1000; i++ {
			counter++
		}
	}

	dec := func() {
		for i := 0; i < 1000; i++ {
			counter--
		}
	}
}
