package main

import "fmt"

type ClassicFunc func(a, b int) int

func GeneralFunc(f ClassicFunc) {
	fmt.Println(f(20, 10))
}

func main() {
	f := func(a, b int) int {
		return a*b - 2
	}
	GeneralFunc(f)
}
