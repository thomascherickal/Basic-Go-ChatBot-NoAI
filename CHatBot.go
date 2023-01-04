package main

import (
	"fmt"
	"strings"
	"testing"
)

func main() {
	fmt.Println(respondToGreeting("hi"))
	fmt.Println(respondToGreeting("hello"))
	fmt.Println(respondToGreeting("bye"))
	fmt.Println(respondToGreeting("goodbye"))
}

func respondToGreeting(greeting string) string {
	greeting = strings.ToLower(greeting)
	if greeting == "hi" || greeting == "hello" {
		return "Hello, User"
	} else if greeting == "bye" || greeting == "goodbye" {
		return "Goodbye, User"
	}
	return "I don't recognize that greeting."
}

func TestRespondToGreeting(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hi", "Hello, User"},
		{"hello", "Hello, User"},
		{"bye", "Goodbye, User"},
		{"goodbye", "Goodbye, User"},
		{"something else", "I don't recognize that greeting."},
	}
	for _, test := range tests {
		if got := respondToGreeting(test.input); got != test.want {
			t.Errorf("respondToGreeting(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}
