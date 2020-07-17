package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var (
		logoString string
		price int = 23
		sum int
		result float32
	)
	fmt.Scan(&logoString)

	sum = utf8.RuneCountInString(logoString) * price

	//fmt.Println((float32(sum)/100))
	//fmt.Println(int32(sum)/100, "р.", )
	//result = string(float32(sum)/100) + "р."
	result = float32(sum)/100
	fmt.Println(result)

	/*if (float32(sum)/100 >= 1) {
		result = string(int32(sum)/100) + "р."
	}*/
}
