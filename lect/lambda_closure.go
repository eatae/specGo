package main

import "fmt"

/* closure */
func Added(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func main() {
	fmt.Println(Added(2)(3))
}
