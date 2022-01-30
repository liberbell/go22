package main

import (
	"fmt"

	m "github.com/rccsdevops/calc/math"
	_ "github.com/rccsdevops/calc/math/stats"
)

var env string

func init() {
	env = "demo"
}

func init() {
	fmt.Println("Database initialized and ready.")
}

func main() {
	fmt.Println(env + " environment loaded.")
	fmt.Println(m.Add(1, 1))
}
