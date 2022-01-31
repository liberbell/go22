package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := inc()
	fmt.Println("x is of type: ", reflect.TypeOf(x))
}

func inc() {

}
