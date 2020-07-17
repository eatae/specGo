package main

import (
	"fmt"
)

func main() {

	// 1
	//var login, email string
	//fmt.Scan(&login)
	//fmt.Scan(&email)
	//
	//if utf8.RuneCountInString(login) < 10 || strings.Contains(login, "@") {
	//	fmt.Println("Некорректный логин")
	//} else if !(strings.Contains(email, "@") && strings.Contains(email, ".")) {
	//	fmt.Println("Некорректная почта")
	//} else {
	//	fmt.Println("ОК")
	//}


	// 2
	const cConst = 75

	age := 26
	var a int
	if fmt.Scan(&a); age*a > 10000 {
		fmt.Println("BOOMER")
	} else {
		fmt.Println("ZOOMER")
	}

}