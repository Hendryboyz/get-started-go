package main

import (
	"fmt"
)

func deferHelloWorld() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

func helloWorld() {
	fmt.Println("World")
	fmt.Println("Hello")
}

func main() {
	helloWorld()
	fmt.Println("---")
	deferHelloWorld()
}
