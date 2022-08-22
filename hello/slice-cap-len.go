package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[0:1]
	printSlice(s)

	// slice bound is determined by the capacity
	s = s[0:3]
	printSlice(s)

	// increase low bound will drop the element
	s = [2:3]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len(s)=%d, cap(s)=%d, %v \n", len(s), cap(s), s)
}
