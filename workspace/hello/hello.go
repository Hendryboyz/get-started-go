package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	hello := "Hello World!"
	fmt.Println(stringutil.Reverse(hello))
	fmt.Println(stringutil.ToUpper(hello))
}
