
package main

import "fmt"

func main() {
	first := 123
	// init
	var firstPtr *int = &first

	fmt.Printf("Type of firstPtr: %T and value %v\n", firstPtr, firstPtr)
	*firstPtr += 200
	fmt.Println("New value of first:", first)

	//Pointer zero value
	var zeroPtr *int
	//*zeroPtr++
	if zeroPtr == nil {
		fmt.Println("Old Value in ZeroPtr:", zeroPtr)
		zeroPtr = &first
		fmt.Println("New Value in ZeroPtr:", zeroPtr)
	}

	// new()
	width := new(int)
	*width++
	fmt.Printf("Type width: %T, value width: %v\n", width, width)
	fmt.Printf("Type of *width %T, and value of *width %v\n", *width, *width)

	// нельзя использовать указатели, нужно разименовывать
	//temp := 36.6
	//tempPtr := &temp
	//tempPtr++

}