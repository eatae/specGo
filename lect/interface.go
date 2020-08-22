package main

import "fmt"

/* INTERFACE */
type Figure interface {
	Perimeter() float64
}

/* Circle */
type Circle struct {
	radius int
}

func (c *Circle) Perimeter() float64 {
	return 3.14 * float64(c.radius) * 2
}

/* Rectangle */
type Rectangle struct {
	a, b int
}

func (r *Rectangle) Perimeter() float64 {
	return float64(2 * (r.a + r.b))
}

/* Triangle */
type Triangle struct {
	a, b, c int
}

func (t *Triangle) Perimeter() float64 {
	return float64(t.a + t.b + t.c)
}

func SummaryPerimeter(figures []Figure) float64 {
	var total float64
	for _, figure := range figures {
		total += figure.Perimeter()
	}
	return total
}

func main() {
	rect := Rectangle{2, 3}
	circle := Circle{5}
	triangle := Triangle{3, 4, 5}
	figures := []Figure{&rect, &circle, &triangle}

	fmt.Println(SummaryPerimeter(figures))
}
