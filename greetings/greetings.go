package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome.", name)
	return message
}

func AnotherHello(name string) string {
	var message string
	message = fmt.Sprintf("Hello, %v. Nice to meet you.", name)
	return message
}
