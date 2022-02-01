package main

import (
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var counter uint16
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				
			}
		}
	}
}
