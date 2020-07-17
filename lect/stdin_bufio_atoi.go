package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	println(input.Scan()) // true

	message, err := strconv.Atoi(input.Text())
	if err != nil {
	// о нет, что-то пошло не так
	}
	fmt.Println(message) // Выводит: 10
}
