package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
	age  int
}

type Employee struct {
	name   string
	id     int
	city   string
	salary int
}

func insertQuery(model interface{}) {
	t := reflect.TypeOf(model)
	v := reflect.ValueOf(model)
	k := t.Kind()
	fmt.Println("Type:", t)
	fmt.Println("Value:", v)
	fmt.Println("Kind:", k)
}

func main() {
	std := Student{"Bob", 21}
	//empl := Employee{"Alice", 903, "Moscow", 100500}
	insertQuery(std)
}
