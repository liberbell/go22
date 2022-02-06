package main

func main() {
	vals := []int{100, 50, 20, 90}
	in := gen(vals)

	fo1 := square(in)
}

func gen(vals []int) <-chan int {
	out := make(chan int)
}
