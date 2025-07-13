package greetings

import (
	"errors"
	"fmt"
)

// A function name starting with a capital letter can be called by a function not in the same package.
// This is known in Go as an exported name.

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		// errors.New returns an error that formats as the given text.
		return "", errors.New("empty name")
	}
	// If a name was given, return a greeting that embeds the name in a message.
	// Sprintf returns a formatted string.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
