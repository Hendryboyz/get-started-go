package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	num := [2]int{0, 1}
	var i int = 0
	return func() int {
		i++
		if i == 1 || i == 2 {
			return i - 1
		}
		tmp := num[1]
		num[1] = num[0] + num[1]
		num[0] = tmp
		return num[1]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
