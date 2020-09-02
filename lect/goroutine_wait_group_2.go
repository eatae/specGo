package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // устанавливаем счетчик в группе на 2
	work := func(id int) {
		defer wg.Done() // уменьшаем счётчик группы на 1
		fmt.Printf("Горутина %d начала выполнение \n", id)
		time.Sleep(1 * time.Second)
		fmt.Printf("Горутина %d завершила выполнение \n", id)
	}
	go work(1)
	go work(2)

	wg.Wait() // ожидаем завершения обоих goroutines
	fmt.Println("END")
}
