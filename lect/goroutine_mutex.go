package main

import (
	"fmt"
	"sync"
)

var counter int = 0 // общий ресурс

func main() {
	ch := make(chan bool) // канал
	var mutex sync.Mutex  // мьютекс
	for i := 1; i < 5; i++ {
		go work(i, ch, &mutex)
	}
	for i := 1; i < 5; i++ {
		<-ch
	}
	fmt.Println("The End")
}

func work(num int, ch chan bool, mutex *sync.Mutex) {
	mutex.Lock() // блокируем доступ к переменной counter
	counter = 0
	for k := 1; k <= 5; k++ {
		counter++
		fmt.Println("Goroutine", num, "-", counter)
	}
	mutex.Unlock() // деблокиуем доступ
	ch <- true
}
