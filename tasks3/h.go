package main

import "fmt"

func main() {
	var (
		count int
		current int
		start int
		end int
		sum int

	)
	fmt.Scan(&count)
	var arr = [25]int{}

	for i := 0; i < count; i++ {
		fmt.Scan(&current)
		arr[i] = current
	}
	fmt.Scan(&start)
	fmt.Scan(&end)

	for i := start-1; i <= end-1; i++ {
		sum += arr[i]
	}

	fmt.Println(sum)
}