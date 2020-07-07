package main

import "fmt"

func main() {
	var (
		word1 = "имя"
		word2 = "твое"
		word3 = "мне"
		word4 = "знакомо"
	)
	fmt.Println(word4, word3, word2, word1)
	fmt.Println(word3, word1, word4, word2)
	fmt.Println(word2, word4, word1, word3)
}