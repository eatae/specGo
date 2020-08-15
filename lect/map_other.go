package main

import "fmt"

func main() {
	employee := initOtherMap()
	fmt.Println(employee)

	iterateOtherMap(employee)
}

// init
func initOtherMap() map[string]int {
	// init
	/*employee := map[string]int{
		"someName": 100,
	}*/
	employee := make(map[string]int)
	employee["Bob"] = 25
	employee["Alice"] = 20
	employee["Petya"] = 23
	if value, ok := employee["Petya"]; ok {
		fmt.Println(value, ok)
	}
	delete(employee, "Petya")

	return employee
}

// iterate
func iterateOtherMap(employee map[string]int) {
	for key, val := range employee {
		fmt.Println("Key:", key, "Value:", val)
	}
}
