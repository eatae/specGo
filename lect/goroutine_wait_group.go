package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker(i int, wg *sync.WaitGroup) {
	fmt.Println("Goroutine #", i, " started!")
	time.Sleep(1 * time.Second)
	fmt.Println("Goroutine #", i, " done!")
	wg.Done() // уменьшаем счётчик группы
}

func main() {
	goNums := 10
	var wg sync.WaitGroup // counter == 0
	wg.Add(goNums)
	for i := 0; i < goNums; i++ {
		go Worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines ended!")
}
