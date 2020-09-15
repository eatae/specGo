package main

import "fmt"

func PrintInfo(name *string, lastName *string) {
	if name == nil {
		panic("runtime error: name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: lastName cannot be nil")
	}
	fmt.Println(*name, *lastName)
	fmt.Println("Done!")
}

func main() {
	name := "Bob"
	PrintInfo(&name, nil)
	fmt.Println("Function main finished")
}
