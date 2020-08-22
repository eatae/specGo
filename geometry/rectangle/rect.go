package rectangle

import (
	"fmt"
	"specGo/geometry/triangle"
)

//A
var A, b = 12, 14

func init() {
	fmt.Println("init rectangle")
	triangle.Show()
}

//Area is function...
func Area(width, length float64) float64 {
	return width * length
}

func Perimeter(width, length float64) float64 {
	return 2 * (width + length)
}
