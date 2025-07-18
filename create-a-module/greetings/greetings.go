package greetings

import (
	"errors"
	"fmt"
	"math/rand"
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
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names,
	// calling the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with the name.
		messages[name] = message
	}
	return messages, nil
}

// randomFormat, an unexported function, randomly returns one of the greeting messages.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.

	// Intn returns a non-negative pseudo-random number in [0,n).
	return formats[rand.Intn(len(formats))]
}
