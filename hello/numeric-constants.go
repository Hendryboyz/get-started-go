package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(i int) int {
	return i*10 + 1
}

func needFloat(f float64) float64 {
	return f * 0.1
}

func main() {
	// too big to print
	// fmt.Printf("%v is Type %T\n", Big, Big)
	fmt.Printf("%v is Type %T\n", Small, Small)
	fmt.Println(needFloat(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needInt(Small))
}
