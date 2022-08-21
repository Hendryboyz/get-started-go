package main

import "fmt"

func main() {
	var i []int = []int{2, 3, 5, 7, 11, 13, 15}
	fmt.Println(i)
	fmt.Println(i[1:3])
	fmt.Println(i[:3])
	fmt.Println(i[1:])
	fmt.Println(i[:])
	
	s := i[1:4]
	fmt.Println(s)
	s = i[:2]
	fmt.Println(s)
	s = i[1:]
	fmt.Println(s)
}