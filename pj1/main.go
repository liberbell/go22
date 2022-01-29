package main

import (
	f "fmt"
	"math"
	. "os"
	_ "strings"

	m "example.local/pj1/pj1/math"
	// "example.local/pj1/pj1/math/stats"
)

func main() {
	// vals := []float64{1, 10, 100, 1000}
	f.Println(m.Add(10, 20))
	// fmt.Println(math.Avg(vals))
	f.Println(math.Pi)
	f.Println(m.Pi)
	f.Fprintln(Stdout, "Hello")
}
