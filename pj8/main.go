package main

func main() {
	for i := 1; i < 5; i++ {
		go myfunc(i)
	}
}
