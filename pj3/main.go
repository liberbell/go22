package main

import (
	"fmt"

	_ "github.com/rccsdevops/calc/math"
	_ "github.com/rccsdevops/calc/math/stats"
)

var env string

func init() {
	env = "demo"
}

func init() {
	fmt.Println("Database initialized and ready.")
}

func init() {
	env = "development"
}

func main() {
	fmt.Println(env + " environment loaded.")
}
