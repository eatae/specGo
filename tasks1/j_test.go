package main

import (
	"fmt"
)

func main() {
	var (
		input1 string
	)
	fmt.Scan(&input1)
	//fmt.Scan(&input2)

	for i := range input1 {
		fmt.Println(input1[i])
		letters := []rune(input1)
		//r, _ := utf8.DecodeRuneInString(input1)

		fmt.Println(len(letters))
	}

	//co, _ := utf8.DecodeRuneInString(int1)
	//res := fmt.Sprintf("%c", co)
	//fmt.Println(res)

	//last1 := int1[len(int1)-1]
	//last2 := (int)(int2[len(int2)-1])
	//foo := utf8.DecodeRuneInString(int1))
	//co, ln := utf8.DecodeRuneInString(int1)
	//fmt.Println(co)
	//fmt.Printf("First rune: %v ln: %v \n", co, ln)

	//fmt.Printf("one: %v", last1)
	//fmt.Printf("two: %v", last2)
	//fmt.Println((last1 + last2) % 2)

	/*if  (last1 + last2) % 2 == 0 {
		fmt.Println("ЧЁТНОЕ")
	} else {
		fmt.Println("НЕЧЁТНОЕ")
	}*/
}
