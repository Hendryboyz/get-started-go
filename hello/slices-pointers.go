package main

import "fmt"

func main() {
	names := [4]string{
		"Henry",
		"Linda",
		"Mary",
		"John",
	}
	fmt.Println(names)
	
	a := names[0:2]
	b := names[2:4]
	fmt.Println(a, b)
	b[0] = "XXX"
	fmt.Println(names)
	fmt.Println(a, b)
}