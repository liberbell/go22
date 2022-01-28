package main

import (
	"fmt"

	"example.local/pj1/pj1/math"
	"example.local/pj1/pj1/math/stats"
)

func main() {
	vals := []float64{1, 10, 100, 1000}
	fmt.Println(math.Add(10, 20))
	fmt.Println(stats.Avg(vals))
}
