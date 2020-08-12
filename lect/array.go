package main

import "fmt"

func main() {

	iteration()
}

func initial() {
	var a [5]int
	fmt.Println("Initial array:", a)
	a[0] = 10
	a[1] = 20
	a[2] = 30
	a[3] = 40
	a[4] = 50
	//a[5] = 2000
	//len(a) - 1 == lastElemIndex
	fmt.Println("After assign:", a)
	//Short-hand array declaration
	b := [3]int{1, 2, 3}
	c := [4]int{200, 300}
	fmt.Println("Arr b:", b)
	fmt.Println("Arr c:", c)

	//Dot-size arr
	dotArr := [...]string{"Bob", "Alice", "Derek"}
	fmt.Printf("%T type and values: %v and length: %v\n", dotArr, dotArr, len(dotArr))

	anotherArr := [4]string{}
	fmt.Printf("%T type and values: %v and length: %v\n", anotherArr, anotherArr, len(anotherArr))
	//anotherArr = dotArr
}

func iteration() {
	var a = [3]string{"Moscow", "SPB", "KZN"}
	b := a
	b[0] = "Monaco"

	//iterating
	for i := 0; i < len(a); i++ {
		fmt.Println("Value:", a[i])
	}

	for i, v := range a {
		fmt.Println("Index:", i, "Value:", v)
	}
	fmt.Println(a)
	fmt.Println(b)
}
