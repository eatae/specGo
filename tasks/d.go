package main

import "fmt"

func main() {
	var (
		name string
		lastName string
		age string

	)
	fmt.Scan(&name)
	fmt.Scan(&lastName)
	fmt.Scan(&age)

	fmt.Println("Имя: ", name," , ", "Фамилия: ", lastName, " , ", "Возраст: ", age, " . ", "Студент BPS")
}
