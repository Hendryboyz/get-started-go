package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	nameError := validaName(name)
	if nameError != nil {
		return "", nameError
	}
	message := fmt.Sprintf("Hi, %v. Welcome.", name)
	return message, nameError
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
