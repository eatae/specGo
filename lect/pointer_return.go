package main

import "fmt"

func changer(num *int) {
	*num += 150
}

func greetings() *int {
	value := 10
	return &value
}

func main() {
	a := 55
	fmt.Println("Old value of a:", a)
	b := &a
	//c := b
	changer(b)
	fmt.Println("New value of a:", a)

	temp := greetings()
	fmt.Println("Result of greetings():", *temp)
}