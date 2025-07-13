// Declare a main package
// Package is a way to group functions, and it's made up of all the files in the same directory.
package main

// Package fmt contains functions for formatting text, including printing to the console.
import (
	"fmt"
	// To do something that might have been implemented by someone else,
	// look for a package that has functions you can use in your code.
	// For instance, this "rsc.io/quote" package.
	// pkg.go.dev site contains many published modules whose packages may have many useful functions.

	// Package quote collects pithy sayings.
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
}

// To run this code.
// $ go run .

// To get a list of other commands.
// $ go help

// By importing packages with external dependencies like "rsc.io/quote",
// add new module requirements and sums
// $ go mod tidy
