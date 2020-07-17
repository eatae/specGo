package main

import "fmt"

func main() {
	var input int

	for {
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		fmt.Println(input)
	}
}