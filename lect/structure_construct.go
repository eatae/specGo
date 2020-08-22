package main

import "fmt"

/* house - не экспортируемая структура, мы можем  */
type house struct {
	number string
	Price  int64
}

/* constructor for house struct */
func New(number string, price int64) house {
	house := house{}
	house.number = number
	house.Price = price
	return house
}

func (house *house) ShowInfo() {
	fmt.Println("house number: ", house.number)
	fmt.Println("house price: ", house.Price)
}

func (house *house) showNumber() {
	fmt.Println("house number: ", house.number)
}

func main() {
	house := New("sss", 12)
	house.showNumber()
}
