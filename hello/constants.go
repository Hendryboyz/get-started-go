package main

import "fmt"

const Pi = 3.14

func Constants() {
	fmt.Println(Pi)
	const World = "世界"
	fmt.Println("Hello", World)
	const Truth = false
	fmt.Println("Go Rules ?", Truth)
}
