package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Go executes init functions automatically at program startup, after global variables have been initialized.
func init() {
	// init sets initial values for variables used in the function.
	rand.Seed(time.Now().UnixNano())
}

// starts with a uppercase letter, making it `public` to code the invoker
func Hello(name string) (string, error) {
	nameError := validaName(name)
	if nameError != nil {
		return "", nameError
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nameError
}

// starts with a lowercase letter, making it accessible only to code in its own package(`private`)
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

func validaName(name string) error {
	// If no name was given, return an error with a message
	if name == "" {
		return errors.New("empty, name")
	} else {
		return nil
	}
}

func AnotherHello(name string) (string, error) {
	nameError := validaName(name)
	if nameError != nil {
		return "", nameError
	}
	var message string
	message = fmt.Sprintf("Hello, %v. Nice to meet you.", name)
	return message, nameError
}
