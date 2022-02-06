package main

import "sync"

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(3)

	go func() {
		wg.Wait()
		close(ch)
	}()

	go func() {
		defer wg.Done()
	}()
}
