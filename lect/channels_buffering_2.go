package main

import "fmt"

func Writer(c chan int) {
	for i := 0; i < 3; i++ {
		c <- i
	}
	fmt.Println("-----------")
	close(c)
}

func main() {
	channel := make(chan int, 3)
	Writer(channel)
	for v := range channel {
		fmt.Println(v)
	}
}
