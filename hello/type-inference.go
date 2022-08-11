package main

import "fmt"

func TypeInference() {
	var i int = 3
	fmt.Printf("%v is type %T\n", i, i)
	j := i
	fmt.Printf("%v is type %T\n", j, j)
	f := 3.14159
	fmt.Printf("%v is type %T\n", f, f)
	g := 0.867 + 0.5i
	fmt.Printf("%v is type %T\n", g, g)
}
