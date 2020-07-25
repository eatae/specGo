package main

import "fmt"

/* Address struct */
type Address struct {
	city   string
	street string
}

/* method */
func (a *Address) PrintFullAddress() {
	fmt.Println("Current city:", a.city, "and street:", a.street)
}

/* Student struct */
type Student struct {
	name string
	age  int
	Address
}

/* method */
func (s *Student) PrintFullInfo() {
	fmt.Println("Student name:", s.name, "and age:", s.age)
}

func main() {
	std := Student{
		name: "Bob",
		age:  21,
		Address: Address{
			city:   "Moscow",
			street: "Baumanskaya",
		},
	}
	std.PrintFullInfo()
	fmt.Println(std.city, std.name)
	std.PrintFullAddress()
}
