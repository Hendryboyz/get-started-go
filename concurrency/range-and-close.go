package main

import "fmt"

func fab(n int, channel chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		channel <- x
		x, y = y, x+y
	}
	close(channel)
}

func main() {
	ch := make(chan int, 10)
	go fab(cap(ch), ch)
	for result := range ch {
		fmt.Println(result)
	}
}
