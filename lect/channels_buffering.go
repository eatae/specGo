package main

import "fmt"

func sendToshiba(c chan string) {
	c <- "Toshiba"
}

func sendLG(c chan string) {
	c <- "LG"
}

func main() {
	channel := make(chan string)
	//channel <- "Sony" 	// значение напрямую
	go sendLG(channel) // значение из goroutine

	fmt.Println(<-channel)
	//fmt.Println(<-channel)
}
