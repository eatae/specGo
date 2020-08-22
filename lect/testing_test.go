package main

import (
	"fmt"
	"testing"
)

type TestTable struct {
	a, b, expected int
}

var foo int64 = 444

var tableAdd = []TestTable{
	{2, 2, 4},
	{3, -3, 0},
	{0, 11, 11},
}

func TestAdd(t *testing.T) {
	for _, item := range tableAdd {
		if item.expected != Add(item.a, item.b) {
			t.Error("Test failed")
		}
	}
}

func main() {
	fmt.Println(tableAdd)
}
