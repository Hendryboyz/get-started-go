package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0) // disable printing the time

	// Get a greeting messages and error
	message, err := greetings.Hello("Henry")
	tryPrint(message, err)

	names := []string{"Mary", "George", "Henry"}
	messages, err2 := greetings.Hellos(names)
	if err2 != nil {
		log.Fatal(err2)
	} else {
		fmt.Println(messages)
	}

	// message2, err3 := greetings.AnotherHello("")
	// tryPrint(message2, err3)
}

func tryPrint(message string, err error) {
	if err != nil {
		log.Fatal(err) // log.Fatal() will print the error and stop the function
	} else {
		fmt.Println(message)
	}
}
