package main

import (
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var counter uint64
	var wg sync.WaitGroup
}
