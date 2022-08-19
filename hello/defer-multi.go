package main

import (
	"fmt"
)

func normalCount() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func deferCount() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
