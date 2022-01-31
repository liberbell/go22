package main

import (
	"fmt"
	"reflect"
)

func main() {
	myfunc := retfunc()
	fmt.Println(myfunc())
	fmt.Println("myfunc is of type:", reflect.TypeOf(myfunc))
}

func retfunc() func() bool {
	return func() bool {
		return true
	}
}
