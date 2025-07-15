package greetings_test

import (
	"regexp"
	"testing"

	"github.com/28604/go-doc-tutorial/create-a-module/greetings"
)

// TestHelloName calls greetings.Hello with a name,
// checking for a valid return value.
func TestHelloName(t *testing.T) {
	// Test functions take a pointer to the testing package's testing.T type as a parameter.
	// You use this parameter's methods for reporting and logging from your test.
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := greetings.Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := greetings.Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want ", error`, msg, err)
	}
}
