package main

import "fmt"

func main() {
	var (
		count int
		sum int
		current int
	)
	fmt.Scan(&count)

	for i := 0; i < count; i++ {
		fmt.Scan(&current)
		if i%2 == 0 {
			sum += current
		} else {
			sum -= current
		}
	}
	fmt.Println(sum)
}