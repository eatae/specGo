package main

import (
	"fmt"
	"time"
)

func Sleeper(chnl chan string) {
	time.Sleep(1500 * time.Millisecond)
	chnl <- "Worker done!"
}

func main() {
	chnl := make(chan string)
	go Sleeper(chnl)
	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case val := <-chnl:
			fmt.Println("Get value:", val)
			return
		default:
			fmt.Println("no value received")
		}
	}
}
