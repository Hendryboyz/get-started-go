package main

import "fmt"

var (
	i int
	s string
	f float64
	b bool
)

func InitialValue() {
	fmt.Printf("%v, %q, %v, %v\n", i, s, f, b)
}
