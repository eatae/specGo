package main

import (
	"bufio"
	"os"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	println(input.Scan()) // true
	println(input.Text()) // STDIN


}
