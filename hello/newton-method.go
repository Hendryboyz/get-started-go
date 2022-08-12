package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := float64(1)
	threshold := math.Pow(10, -15)

	var nextZ float64
	var i int

	for {
		// newton's method
		nextZ = (z - (z*z-x)/(2*z))
		if math.Abs(nextZ-z) <= threshold {
			break
		}
		z = nextZ
		i++
	}
	fmt.Printf("iteration stop at %v\n", i)
	return z
}

func main() {
	fmt.Println(sqrt(7))
}
