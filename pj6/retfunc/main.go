package main

import "fmt"

func main() {
	myfunc := retfunc()
	fmt.Println(myfunc())
}

func retfunc() func() bool {
	return func() bool {
		return true
	}
}
