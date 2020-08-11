package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func main() {
	// name
	scanner.Scan()
	name := scanner.Text()
	// last name
	scanner.Scan()
	lastName := scanner.Text()
	// age
	scanner.Scan()
	age := scanner.Text()

	fmt.Printf("Имя: %v, Фамилия: %v, Возраст: %v. Студент BPS", name, lastName, age)
}
