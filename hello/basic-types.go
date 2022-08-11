package main

import (
	"fmt"
	"math/cmplx"
)

var (
	c            bool       = false
	pythonString string     = "python"
	thatFloat    float64    = 3.14159
	thisComplex  complex128 = cmplx.Sqrt(-5 + 12i)
)

// Not allow to initial value in global
// myInt := 3

func PrintTypes() {
	fmt.Printf("Type %T and value: %v\n", c, c)
	fmt.Printf("Type %T and value: %v\n", pythonString, pythonString)
	fmt.Printf("Type %T and value: %v\n", thatFloat, thatFloat)
	fmt.Printf("Type %T and value: %v\n", thisComplex, thisComplex)
}
