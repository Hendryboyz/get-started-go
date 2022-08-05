// Ending a file's name with _test.go tells the `go test` to test it
package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Henry"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := Hello("Henry")
	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("Henry") = %q, %v, want match for %#q, nil`, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, err)
	}
}
