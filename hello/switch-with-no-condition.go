package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("Current time: %s\n", now)
	switch {
	case now.Hour() < 12:
		fmt.Println("Good Morning")
	case now.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}
}
