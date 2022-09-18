package main

import "fmt"

func fab(n, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case <-quit:
			fmt.Println("quit")
			return
		case n <- x:
			x, y = y, x+y
		}
	}
}

func main() {
	n := make(chan int, 10)
	quit := make(chan int)
	go func() {
		for i := 0; i < cap(n); i++ {
			fmt.Println(<-n)
		}
		quit <- 0
	}()
	fab(n, quit)
}
