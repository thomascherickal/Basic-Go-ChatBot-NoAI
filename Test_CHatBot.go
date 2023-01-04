package main

import "testing"

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
		if got := TestRespondToGreeting(test.input); got != test.want {
			t.Errorf("respondToGreeting(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}
