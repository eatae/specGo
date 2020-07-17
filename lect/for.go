package main

import "fmt"

func main() {

	//classic for loop and break /continue
	for i := 0; i <= 10; i++ {
		if i == 6 {
			break
		}

		if i == 4 {
			continue
		}
		fmt.Printf("%v and squared %v\n", i, i*i)
	}
	fmt.Println("FINISH")


	//nested loops
outer:
	for i := 1; i < 10; i++ {
		for j := 3; j < 14; j++ {
			if i == j {
				break outer
			}
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	}

	age := 21
outer2:
	for i := 0; i < 10; i++ {
		switch age {
		case 20:
			fmt.Println("1")
		case 21:
			fmt.Println("2")
			break outer2
		default:
			fmt.Println("3")
		}
	}

}