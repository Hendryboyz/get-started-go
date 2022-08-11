package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	fmt.Printf("%T %T is %v, %v\n", x, y, x, y)
	var float float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Printf("%T is %v\n", float, float)
	z := uint(float)
	fmt.Printf("%v, %v, %v\n", x, y, z)
}
