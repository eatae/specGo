package main

import (
	"specProject/lect/str"
	"fmt"
)

func PrintEmployee(empl str.Employee) {
	fmt.Println("Name:", empl.FirstName)
	fmt.Println("Lastname:", empl.LastName)
	fmt.Println("Age:", empl.Age)
}

func main() {
	empl1 := str.Employee{
		FirstName: "Bob",
		LastName:  "Snow",
		Age:       27,
	}

	empl2 := str.Employee{
		FirstName: "Bob",
		LastName:  "Snow",
		Age:       28,
	}

	if empl2 == empl1 {
		fmt.Println("EQUALS")
	}

	fmt.Println("-----------")
	PrintEmployee(empl2)
}