package main

import "fmt"

/* Employee */
type Employee struct {
	firstName string
	lastName  string
	age       int
}

/* Structure */
type Student struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	/* student */
	student1 := Student{"Petr", "Ivanov", 24}

	fmt.Println(student1)
}
