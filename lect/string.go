package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var (
		name     string = "John"
		lastName string = "Иванов"
	)
	fmt.Println("Concatenate:", name+" "+lastName)
	fmt.Println(len("John")) // 4
	fmt.Println(len([]int{1, 2, 3}))
	fmt.Println(len("Анна"))                    // кириллица 8 Байт
	fmt.Println(utf8.RuneCountInString("Анна")) // кириллица 4 символа
}
