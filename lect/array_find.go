package main

import "fmt"

func main() {
	var names = [3]string{"Bob", "Alice", "Nicolay"}
	searched := "Fedya"
	fmt.Println(Find(names, searched))
}
