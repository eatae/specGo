package main

import (
	"fmt"
)

func reverseMessage(message string) {
	defer fmt.Println()
	for _, char := range message {
		defer fmt.Printf("%c", char) // reverse
	}
}

func main() {
	message := "Hello world!"
	reverseMessage(message)
}
