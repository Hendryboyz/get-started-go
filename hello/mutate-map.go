package main

import "fmt"

func main() {
	m := make(map[string]int)
	
	m["ans"] = 42
	fmt.Println("The value is", m["ans"])
	
	m["ans"] = 48
	fmt.Println("The value is", m["ans"])
	
	delete(m, "ans")
	fmt.Println("The value is", m["ans"])
	
	element, isExist := m["ans"]
	fmt.Println("The value", element, ", Present ?", isExist)
}