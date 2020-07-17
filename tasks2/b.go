package main

import "fmt"

func main() {
	//Ax^2 + Bx + C = 0
	var A, B, C float64
	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)

	if A != 0 {
		D := B*B - 4*A*C
		if D > 0 {
			fmt.Println("два корня")
		} else if D == 0 {
			fmt.Println("один корень")
		} else {
			fmt.Println("корней нет")
		}
	} else {
		//Bx + C = 0
		// x = -C/B
		if B == 0 {
			fmt.Println("корней нет")
		} else {
			fmt.Println("один корень")
		}
	}
}