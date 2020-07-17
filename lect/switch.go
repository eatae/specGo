package main

import "fmt"

func main() {

	switch age := 31; age {
	case 28, 29, 30:
		fmt.Println("Age is 28-30")
		fallthrough
	case 31, 32, 33:
		fmt.Println("Age is 31-33")
		fallthrough
	// case 29:
	// 	fmt.Println()
	case 34, 35, 40:
		fmt.Println("Age is 34-40")
	default:
		fmt.Println("?")
	}
}