package main

import "fmt"

func main() {
	var (
		one string
		two string
		three string
	)
	fmt.Scan(&one)
	fmt.Scan(&two)
	fmt.Scan(&three)

	fmt.Println(three+two+one)
}
