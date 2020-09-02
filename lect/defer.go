package main

import (
	"fmt"
	"time"
)

func testFunc() {
	fmt.Println("Finish testFunc")
}

func main() {
	defer testFunc()
	fmt.Println("Main function started!")
	time.Sleep(2 * time.Second)
	fmt.Println("Main function finished")
}
