package main

import (
	"fmt"
	// The way to import greetings if we actually want to download from GitHub.
	"github.com/28604/go-doc-tutorial/create-a-module/greetings"
)

func main() {
	message := greetings.Hello("Chen")
	fmt.Println(message)
}

// However, if we want to find the greetings package locally, we run
// $ go mod edit -replace github.com/28604/go-doc-tutorial/create-a-module/greetings=../greetings
// $ go mod tidy
