package concurrency

import "fmt"

func main() {
	bch := make(chan int)
	bch <- 1
	bch <- 2
	fmt.Println(<-bch)
	fmt.Println(<-bch)
}
