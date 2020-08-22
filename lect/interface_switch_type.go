package main

import (
	"fmt"
	"reflect"
)

type Worked interface {
	Work()
}

type Person struct {
	name string
	age  int
}

func (p Person) Work() {
	fmt.Println("Person is Worker")
}

func FindType(i interface{}) {
	switch t := i.(type) {
	case string:
		fmt.Println("This is string and value: ", i.(string))
	case int:
		fmt.Println("This is int and value: ", i.(int))
	case float64:
		fmt.Println("This is float64 and value: ", i.(float64))
	case Worked:
		t.Work()
	default:
		fmt.Println("Unknown type")
	}
}

func CheckTypeByReflection(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
}

func main() {
	FindType("Petr")
	CheckTypeByReflection(14.5)
}
