package main

import "fmt"

// 143  ->  ((1*1)+(4*4)+(3*3)) + ((1*1*1)+(4*4*4)+(3*3*3))

func calcSquares(num int, sqr_channel chan int) {
	summ := 0
	for num != 0 {
		d := num % 10
		summ += d * d
		num /= 10
	}
	sqr_channel <- summ
}

func calcCubes(num int, cube_channel chan int) {
	summ := 0
	for num != 0 {
		d := num % 10
		summ += d * d * d
		num /= 10
	}
	cube_channel <- summ
}

func main() {
	num := 143
	sqr_channel := make(chan int)
	cube_channel := make(chan int)

	go calcSquares(num, sqr_channel)
	go calcCubes(num, cube_channel)

	square, cube := <-sqr_channel, <-cube_channel
	result := square + cube
	fmt.Println("Result is: ", result)
}
