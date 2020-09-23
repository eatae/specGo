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

func iterateStruct(s interface{}) {
	if reflect.TypeOf(s).Kind() == reflect.Struct {
		v := reflect.ValueOf(s)
		countFields := v.NumField()
		for i := 0; i < countFields; i++ {
			fmt.Printf("Field: %d Type: %v Value: %v\n", i, v.Field(i).Type(), v.Field(i))
		}
	}
}

func main() {
	std := Student{"Bob", 21}
	iterateStruct(std)
}
