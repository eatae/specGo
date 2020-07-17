package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		count int
		task string
		quantity int
	)
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	strconv.ParseInt(input.Text(), 10, 64)
	var arr = [25]string{}
	var result = [25]string{}

	for i := 0; i < count; i++ {
		fmt.Scan(&task)
		arr[i] = task
	}
	fmt.Scan(&quantity)





	input := bufio.NewScanner(os.Stdin)

	for i := 0; i < len(words); i++ {
		var message string
		input.Scan()
		message = strconv.ParseFloat(input.Text(), 64)
		words[i] = message
	}
}