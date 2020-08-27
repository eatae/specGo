package main

import "fmt"

func numbers(chNums chan int) {
	for i := 0; i < 50; i++ {
		chNums <- i
	}
	close(chNums) // закрываем канал
}

func main() {
	chNums := make(chan int)
	go numbers(chNums)
	//for {
	//	if val, ok := <- chNums; !ok {
	//		break						// если закрыт канал - выходим из цикла
	//	} else {
	//		fmt.Println("Receive from channel: ", val)
	//	}
	//}
	// или
	for val := range chNums { // пока итерируется канал
		fmt.Println("Receive from channel: ", val)
	}
}
