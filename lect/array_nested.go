package main

import "fmt"

func main() {
	words := [3][2]string{
		{"a", "b"},
		{"c", "d"},
		{"e", "f"},
	}

	for _, v1 := range words {
		for _, v2 := range v1 {
			fmt.Printf("%v ", v2)
		}
		fmt.Printf("\n")
	}
}