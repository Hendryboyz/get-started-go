package main

import "fmt"

func main() {
	forSum := 0
	for i := 1; i <= 10; i++ {
		forSum += i
	}
	fmt.Println(forSum)

	whileSum := 1
	for whileSum < 1000 {
		whileSum += whileSum
	}
	fmt.Println(whileSum)
}
