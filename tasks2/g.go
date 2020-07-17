package main

import (
	"fmt"
)


var (
	min float32 = 100.0
	max float32 = 140.0
	pulse float32
	count int
	minItem float32 = 140.0
	maxItem float32 = 100.0
)

func main() {
	for pulse >= 0 {
		fmt.Scan(&pulse)
		if pulse >= min && pulse <= max {
			count ++
			if pulse < minItem {
				minItem = pulse
			}
			if pulse > maxItem {
				maxItem = pulse
			}
		}
	}
	fmt.Println(count)
	fmt.Println(minItem, maxItem)
}