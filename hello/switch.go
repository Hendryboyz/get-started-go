package main

import (
	"fmt"
	"runtime"
	"time"
)

func whenSaturday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Println("Today lol")
	case today + 1:
		fmt.Println("Tomorrow OTZ")
	case today + 2:
		fmt.Println("In 2 days QQ")
	default:
		fmt.Println("Too far to come")
	}
}

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("This is OS X")
	case "linux":
		fmt.Println("This is Linux")
	default:
		fmt.Printf("%s. \n", os)
	}
	whenSaturday()
}
