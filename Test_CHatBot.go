package NoAI

import (
	"fmt"
	"strings"
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
