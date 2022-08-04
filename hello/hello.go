package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0) // disable printing the time

	// Get a greeting message and error
	message, err := greetings.Hello("Henry")
	tryPrint(message, err)

	message2, err2 := greetings.AnotherHello("")
	tryPrint(message2, err2)
}

func tryPrint(message string, err error) {
	if err != nil {
		log.Fatal(err) // log.Fatal() will print the error and stop the function
	} else {
		fmt.Println(message)
	}
}
