package main

import (
	"fmt"
	"math"
)

func sqrt(f float64) string {
	if f < 0 {
		return sqrt(-f) + "i"
	}
	return fmt.Sprint(math.Sqrt(f))
}

func pow(x, n, limit float64) float64 {
	if value := math.Pow(x, n); value < limit {
		return value
	} else {
		fmt.Printf("%g >= %g\n", value, limit)
	}
	return limit
}

func main() {
	fmt.Println(sqrt(4), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
