package main

import "fmt"

func someFunc(arg int) {
	fmt.Println("someFunc argument = ", arg)
}

func main() {
	num := 10
	defer someFunc(num)
	num = 30
	fmt.Println("Finished main func with num = ", num)
}
