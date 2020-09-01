package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker(i int, wg *sync.WaitGroup) {
	fmt.Println("Goroutine #", i, "Started!")
	time.Sleep(2 * time.Second)
	fmt.Println("Goroutine")
}
