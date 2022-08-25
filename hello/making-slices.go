package main

import "fmt"

func printSlice(name string, s []int) {
	fmt.Printf("%s: len(s)=%d, cap(s)=%d, %v\n", name, len(s), cap(s), s)
}

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:3]
	printSlice("c", c)

	d := c[2:3]
	printSlice("d", d) // cap(d) = 3, len(d) = 1
}
