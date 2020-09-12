package main

import (
	"fmt"
	"sync"
)

type Rectangle struct {
	a, b int
}

func (r Rectangle) Area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.a < 0 {
		fmt.Println("Side a less the zero!")
		return
	}
	if r.b < 0 {
		fmt.Println("Side b less the zero!")
		return
	}
	area := r.a * r.b
	fmt.Println("Area of ", r, "is:", area)
}

func main() {
	var wg sync.WaitGroup
	r1 := Rectangle{-3, 10}
	r2 := Rectangle{4, -1}
	r3 := Rectangle{20, 10}
	r4 := Rectangle{10, 10}
	r5 := Rectangle{1, 12}
	rectSlice := []Rectangle{r1, r2, r3, r4, r5}
	for _, v := range rectSlice {
		wg.Add(1)
		go v.Area(&wg)
	}
	wg.Wait()
	fmt.Println("All Goroutines done!")
}
