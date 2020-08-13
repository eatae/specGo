package main

import "fmt"

func trianglePerimeter(a, b, c float64) (result, area float64) {
	result = a + b + c
	area = a * b * c // WRONG
	return
}

func Area(w, l float64) float64 {
	result := w * l
	return result
}

func Perimeter(w float64, l float64) float64 {
	result := 2 * (w + l)
	return result
}

func geometryMetrics(w, l float64) (float64, float64) {
	return Area(w, l), Perimeter(w, l)
}

// variadic parameter a ...int
func calcSum(a ...int) int {
	fmt.Printf("%T type and %v values\n", a, a)
	return 2
}

// FindNum for ...
func FindNum(num int, nums ...int) {
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "was found with id:", i)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in", nums)
	}
}

// FindNum(20, slice...)

func main() {
	width, length := 23.5, 12.5
	area := Area(width, length)
	perimeter := Perimeter(width, length)

	fmt.Println("Area:", area)
	fmt.Println("Perimeter:", perimeter)

	_, newPerimeter := geometryMetrics(12, 12)
	//fmt.Println("New area is:", newArea)
	fmt.Println("New perimeter is:", newPerimeter)

	aTriangle, pTriangle := trianglePerimeter(10, 20, 30)
	fmt.Println("Area triangle:", aTriangle)
	fmt.Println("Perimeter triangle:", pTriangle)

	// Lambda function
	temp := func(a, b int) int {
		return a*a + 2*a*b + b*b
	}(10, 20)

	// Lambda function
	lambdaFunction := func(a, b int) int {
		return a*a + 2*a*b + b*b
	}

	fmt.Println("Result of lambda expression:", temp)
	fmt.Printf("%T type of lambda function\n", lambdaFunction)
	fmt.Printf("%T type for Area\n", Area)

	lambdaFunction(20, 30)

	slice := []int{20, 30, 40}
	FindNum(20, slice...)

}
