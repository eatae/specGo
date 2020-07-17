package main

import "fmt"

func main() {
	var (
		width int
		height int
	)
	fmt.Scan(&width)
	fmt.Scan(&height)

	fmt.Println("Периметр:",(width+height)*2)
	fmt.Println("Площадь:", width*height)
}
