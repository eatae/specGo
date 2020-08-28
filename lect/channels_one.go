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
	var result int
	share_channel := make(chan int) // один канал для пары goroutine

	go calcSquares(num, share_channel)
	go calcCubes(num, share_channel)
	result += <-share_channel // получили первое значение
	result += <-share_channel // получили второе значение

	fmt.Println("Result is: ", result)
}
