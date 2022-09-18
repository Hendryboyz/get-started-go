package concurrency

import (
	"fmt"
	"time"
)

func say(sentence string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(sentence)
	}
}

func main() {
	go say("Hello")
	say("Hello World")
}
