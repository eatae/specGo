package main

import "fmt"


func main() {
	employee := initMap()
	iterateMap(employee)
}



// init
func initMap() map[string]int {
	// 1
	map1 := make(map[string]string)
	map1["one"] = "1"
	map1["one"] = "2"
	// refer type
	map2 := map1
	map2["one"] = "eins"
	// show
	//fmt.Println("Map1:", map1)
	//fmt.Println("Map2:", map2)

	// можно сравнивать только с nil
	if map1 == nil {
		fmt.Println(map1)
	}

	// 2
	employee := map[string]int{
		"Steve": 43,
		"Alex":  66,
	}
	employee["Petya"] = 33
	if value, ok := employee["Petya"]; ok {
		fmt.Println(value)
	}

	return employee
}

// iterate
func iterateMap(employee map[string]int) {
	for key, val := range employee {
		fmt.Println("Key:", key, "Value:", val)
	}
}