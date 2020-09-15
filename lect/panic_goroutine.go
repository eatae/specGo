package main

import "fmt"

func RecoveryDivisionGoroutine(done chan bool) {
	if rec := recover(); rec != nil {
		fmt.Println("Recovered", rec)
	}
	done <- false
}

func Calculator(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	done := make(chan bool)
	go Division(a, b, done)
	<-done
}

func Division(a, b int, done chan bool) {
	defer RecoveryDivisionGoroutine(done) // эта паника должна отлавливаться здесь
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	done <- true
}

func main() {
	Calculator(440, 0)
	fmt.Println("Main function Done")
}
