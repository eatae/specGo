package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	nameDates := [][]string{}
	qtyFriends, _ := strconv.Atoi(getStdin())
	for i := 0; i < qtyFriends; i++ {
		nameDates[i] = strings.Fields(getStdin())
	}

	months := []string{}
	qtyMonths, _ := strconv.Atoi(getStdin())
	for i := 0; i < qtyMonths; i++ {
		months[i] = getStdin()
	}

}

func getStdin() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}