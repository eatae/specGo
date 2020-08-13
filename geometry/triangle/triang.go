package triangle

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("init triangle")
}

//Show
func Show() {
	fmt.Println("Show func fo example")
}

//Perimeter
func Perimeter(a, b, c float64) float64 {
	return a + b + c
}

//Area
func Area(a, b, c float64) float64 {
	p := Perimeter(a, b, c) / 2
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}
