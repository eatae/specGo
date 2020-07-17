package main

import "fmt"

func main() {
	var (
		word1 string
		word2 string
		word3 string
	)
	fmt.Scan(&word1)
	fmt.Scan(&word2)
	fmt.Scan(&word3)

	if (word1 == "раз" || word1 == "один") && word2 == "два" && word3 == "три" {
		fmt.Println("ОК")
	} else if word1 == "1" && word2 == "2" && word3 == "3" {
		fmt.Println("ОК")
	} else {
		fmt.Println("НЕ ПРАВИЛЬНО")
	}
}
