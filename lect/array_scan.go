package main

import "fmt"

func main() {
	var nums [5]int
	var sum int

	for i := 0; i < len(nums); i++ {
		var tmp int
		fmt.Scan(&tmp)
		nums[i] = tmp
		sum += tmp
	}

	fmt.Println(nums)
	fmt.Println(sum)
}
