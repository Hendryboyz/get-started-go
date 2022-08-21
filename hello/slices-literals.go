package main

import "fmt"

func main() {
	i := []int{2, 3, 5, 8, 13, 21}
	fmt.Println(i)
	
	s := []struct{
		i int
		b bool
	} {
		{1, true},
		{2, false},
		{3, true},
		{4, false},
		{5, true},
	}
	fmt.Println(s)
	
	b := []bool{true, false, false, true, true, false}
	fmt.Println(b)
}