package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		input1 string
		input2 string
	)
	fmt.Scan(&input1)
	fmt.Scan(&input2)

	runes1 := []rune(input1)
	runes2 := []rune(input2)
	lastLetter1 := string(runes1[len(runes1)-1])
	lastLetter2 := string(runes2[len(runes2)-1])
	lastInt1, _ := strconv.Atoi(lastLetter1)
	lastInt2, _ := strconv.Atoi(lastLetter2)

	if (lastInt1+lastInt2)%2 == 0 {
		fmt.Println("ЧЁТНОЕ")
	} else {
		fmt.Println("НЕЧЁТНОЕ")
	}
}
