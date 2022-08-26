package main

import "fmt"

func main() {
	pow := make([]int, 11)
	fmt.Printf("cap(pow) = %d, len(pow) = %d\n", cap(pow), len(pow))
	for i := range pow {
		pow[i] = 1 << i
	}

	for _, value := range pow {
		fmt.Printf("%d ", value)
	}
}
