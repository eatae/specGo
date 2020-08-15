package main

import (
	"fmt"
	"specGO/geometry/rectangle"
)

func init() {
	fmt.Println("Init geometry")
}

func main() {
	fmt.Println("Geometry start here")
	w, l := 22.5, 43.2
	area := rectangle.Area(w, l)
	perimeter := rectangle.Perimeter(w, l)

	fmt.Println("Area: ", area)
	fmt.Println("Perimeter: ", perimeter)
	fmt.Println("A: ", rectangle.A)

	slice := []int{0, 10, 20, 30}
	fmt.Println(slice)
	Modify(slice)
	fmt.Println(slice)
}

func Modify(sl []int) {
	sl[0] = 100
	sl = append(sl, 2000)
}
