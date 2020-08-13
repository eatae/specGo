package main

import (
	"fmt"
	"sort"
)

func main() {
	//initialSlice()
	sortSlice()
}

func initialSlice() {

	arr := [5]int{10, 20, 30, 40, 50}

	// by array
	var slice1 []int = arr[1:4]
	fmt.Printf("slice1 type: %T || value: %v\n", slice1, slice1)

	// slice
	slice2 := []int{1, 2, 3}
	fmt.Println("slice2:", slice2)

	// make
	slice3 := make([]int, 0, 6)
	fmt.Println("Slice3:", slice3)

	// copy value, not referal
	slice4 := append([]int{}, slice2...)
	fmt.Println("slice4:", slice4)
	slice4[0] = 100
	fmt.Println("slice4:", slice4)
	fmt.Println("slice2:", slice2)

	// copy value, not referal
	city := []string{"Moscow", "SPB", "KZN"}
	newCity := make([]string, len(city)-1)
	// copy
	copy(newCity, city)
	newCity[0] = "Monaco"
	fmt.Println("New city:", newCity[:2])
	fmt.Println("City:", city)

}

func iterationSlice() {

	slice := []int{1, 2, 3}

	// for by range
	for i := range slice {
		slice[i] += 200
	}
}

func sortSlice() {
	var a = []int{300, 400, 500, -1, 20, 40}
	sort.Ints(a)
	fmt.Println(a)
}
