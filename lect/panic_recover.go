package main

import (
	"fmt"
)

func ShowInfo(name *string, lastName *string) {
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
	/* recover - отлавливаем panic */
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("recovered from ", rec)
			//debug.PrintStack()
		}
	}()
	name := "Bob"
	ShowInfo(&name, nil)
	fmt.Println("Function main finished")
}
