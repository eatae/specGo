package main

import "fmt"

func main() {
	var (
		word1 string
		word2 string
		word3 string
		word4 string
		name string
	)
	fmt.Scan(&word1)
	fmt.Scan(&word2)
	fmt.Scan(&word3)
	fmt.Scan(&word4)
	fmt.Scan(&name)

	fmt.Printf("%s - %s\n", word4, name)
	fmt.Printf("%s - %s\n", word3, name)
	fmt.Printf("%s - %s\n", word2, name)
	fmt.Printf("%s - %s\n", word1, name)
}
