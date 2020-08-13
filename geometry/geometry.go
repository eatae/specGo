package main

import (
	"fmt"
	"specGO/geometry/rectangle"
)

func init() {
	testFunc()
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
}

func testFunc() {
	fmt.Println("Test func!")
}
