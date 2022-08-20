package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	var p *Vertex
	p = &v
	fmt.Println(*p)
	p.X = 1e9
	fmt.Println(*p)
}
