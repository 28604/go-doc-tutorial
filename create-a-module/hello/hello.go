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

	// Request a greeting message.
	message, err := greetings.Hello("Chen")
	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message to the console.
	fmt.Println(message)
}

// However, if we want to find the greetings package locally, we run
// $ go mod edit -replace github.com/28604/go-doc-tutorial/create-a-module/greetings=../greetings
// $ go mod tidy
