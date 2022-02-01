package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i < 5; i++ {
		go myfunc(i)
	}
	time.Sleep(10 * time.Millisecond)
	fmt.Println(("Each goroutine has run to competion, thanks for waiting."))
}

func myfunc(i int) {
	fmt.Println("Finished executing iteration.", i)
}
