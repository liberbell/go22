package math

import "fmt"

// This is a comment.
const Pi = 3.14159

func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

func init() {
	fmt.Println("initialized package math (local).")
}
