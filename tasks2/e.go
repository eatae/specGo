package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		pass1 string
		pass2 string
	)
	for {
		fmt.Scan(&pass1)
		fmt.Scan(&pass2)

		//fmt.Println(len(pass1))
		//fmt.Println(strings.Contains(pass1, "qwe"))

		if len(pass1) < 8 {
			fmt.Println("Слишком короткий пароль!")
			continue
		} else if strings.Contains(pass1, "qwe") || strings.Contains(pass1, "123") {
			fmt.Println("Слишком простой пароль!")
			continue
		} else if pass1 != pass2 {
			fmt.Println("Введенные пароли различаются!")
			continue
		} else {
			fmt.Println("Ну наконец-то!")
			break
		}
	}
}