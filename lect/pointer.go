package main

import "fmt"

func main() {
	var first int = 123
	var firstPointer *int = &first
	fmt.Println(firstPointer)

	/* pointer zero value == nil */
	var zeroPtr *int
	fmt.Println(zeroPtr)

	/* разыменование */
	*firstPointer = 200
	fmt.Println(first)

	/* default type pointer by new() */
	defaultPointer := new(int) // create pointer to default int value
	*defaultPointer++
	fmt.Printf("Type defaultPointer: %T, value defaultPointer: %v\n", defaultPointer, defaultPointer)
	fmt.Printf("Type *defaultPointer: %T, value *defaultPointer: %v\n", *defaultPointer, *defaultPointer)

	/* pointer in function args  */
	a := 55
	aPointer := &a
	changer(aPointer)
	fmt.Println(a)

	/* return a pointer from a function */
	temp := returning() // вернётся указатель
	fmt.Println(*temp)

	/* go way ideomatic */
	nums := [3]int{1, 2, 3}
	sliceChanger(nums[:])
	fmt.Println(nums)

}

/* function expected pointer */
func changer(num *int) {
	*num += 10
}

/* function return a pointer from */
func returning() *int {
	someVal := 10
	return &someVal // возвращаем ссылку на 10
}

/* go way ideomatic */
func sliceChanger(sls []int) {
	sls[0] = 500
}
