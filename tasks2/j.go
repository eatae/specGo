package main

import (
	"fmt"
)

func main() {

	var pos int
	fmt.Scan(&pos)

	for neg := pos*(-1); neg <= pos; neg++{
		fmt.Println("Квадрат числа", neg, "равен", neg*neg)
	}
}