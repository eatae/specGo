package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		total string = "чат"
		message string
	)
	fmt.Scan(&message)

	if strings.Contains(message, total) {
		fmt.Println("БОТ")
	} else {
		fmt.Println("НЕ БОТ")
	}
}