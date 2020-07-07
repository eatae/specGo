package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var (
		login string
		mail string
	)
	fmt.Scan(&login)
	fmt.Scan(&mail)

	if strings.Contains(login, "@") || utf8.RuneCountInString(login) < 10 {
		fmt.Println("Некорректный логин")
	} else if !strings.Contains(mail, "@") || !strings.Contains(mail, ".") {
		fmt.Println("Некорректная почта")
	} else {
		fmt.Println("ОК")
	}
}
