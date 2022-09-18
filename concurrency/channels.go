package concurrency

import "fmt"

func sum(nums []int, c chan int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	c <- sum
}

func main() {
	nums := []int{4, 3, 7, 9, 2, 1}
	fmt.Println(nums)
	c := make(chan int)
	go sum(nums[:len(nums)/2], c)
	go sum(nums[len(nums)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
