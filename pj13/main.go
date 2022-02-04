package main

func main() {
	ch := make(chan int)
	mult := func(x y, int)  {
		res := x * y
		ch <- res
	}
	val := <-ch
	fmt.Printf()
}
