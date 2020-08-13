package circle

import (
	_ "fmt"
	"math"
)

func Area(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}
