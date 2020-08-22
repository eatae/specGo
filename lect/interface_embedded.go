package main

import (
	"fmt"
	"reflect"
)

type Duck interface {
	Quack()
}
type SomeDuck struct {
	name string
}

func (d SomeDuck) Quack() {
	fmt.Println("Person is Worker")
}

type Dog interface {
	Yelp()
}

/* embedded interface */
type Combo interface {
	Duck
	Dog
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
