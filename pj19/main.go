package main

import "sync"

func main() {
	vals := []int{100, 50, 20, 90}
	in := gen(vals)

	fo1 := square(in)
}

func gen(vals []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, val := range vals {
			out <- val
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for val := range in {
			out <- val * val
		}
		close(out)
	}()
	return out
}

func merge(fo ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	fi := func(ch <-chan int) {
		for val := range ch {
			out <- val
		}
	}
}
