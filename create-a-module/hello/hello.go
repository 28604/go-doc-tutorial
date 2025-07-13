package main

import (
	"fmt"
	"log"

	"github.com/28604/go-doc-tutorial/create-a-module/greetings"
)

func main() {
	// Set prefix of the logger.
	log.SetPrefix("greetings: ")
	// Disable printing time, source file, and line number of the logger.
	log.SetFlags(0)

	// Request greeting messages for the names.
	names := []string{"John", "Jack", "James"}
	messages, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned map of messages to the console.
	fmt.Println(messages)
}

// However, if we want to find the greetings package locally, we run
// $ go mod edit -replace github.com/28604/go-doc-tutorial/create-a-module/greetings=../greetings
// $ go mod tidy
