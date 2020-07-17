package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var (
		message string
	)
	fmt.Scan(&message)
	price := utf8.RuneCountInString(message)*23
	if price / 100 < 1 {
		fmt.Println(price, "коп.")
	} else {
		fmt.Println(price/100, "р.", price%100, "коп.")
	}
}
