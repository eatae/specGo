package main

import (
	"fmt"
	"time"
)

func main() {
	go func(name string) {
		fmt.Println("Hello", name)
	}("Alex")
	time.Sleep(1 * time.Second)
}
