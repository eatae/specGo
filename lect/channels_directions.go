package main

import (
	"fmt"
	"time"
)

func pingerD(c chan<- string) { // принимающий канал
	for i := 0; ; i++ {
		c <- "ping" // направляем значение в канал
	}
}

func pongerD(c chan<- string) { // принимающий канал
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printerD(c <-chan string) {
	for {
		msg := <-c // получаем значение из канала
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string) // создаём двунаправленный строковый канал

	go pingerD(c)  // 1. запись строки в канал
	go pongerD(c)  // 3. запись строки в канал
	go printerD(c) // 2. 4. получение значения из канала

	var input string
	fmt.Scanln(&input) // выводится по очереди ping/pong
}
