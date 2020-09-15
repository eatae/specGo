package main

import "fmt"

type Addition func(a, b int) int

func main() {
	var a Addition = func(a, b int) int {
		return a + b
	}
	res := a(10, 20)
	fmt.Println("Result is", res)
	fmt.Printf("Type: %T \n Value: %T\n", a, a)
}
