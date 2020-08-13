package main

import (
	"fmt"
	"sort"
)

func main() {
	/* slice into slice */
	numbers := []int{1, 2, 3, 4, 5}
	newNumbers := []int{10, 20, 30, 40, 50}
	numbers = append(numbers, newNumbers...)
	//fmt.Println(numbers)

	/* multi slice */
	langs := [][]string{
		{"c++", "c", "RUST"},
		{"JS", "Python", "PHP"},
	}
	fmt.Print(langs)

	/* find */

}

func sortSlice() {
	var a = []int{300, 400, 500, -1, 20, 40}
	sort.Ints(a)
	fmt.Println(a)
}
