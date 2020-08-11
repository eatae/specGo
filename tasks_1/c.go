package main

import "fmt"

var (
	name string = "имя"
	your string = "твое"
	i    string = "мне"
	know string = "знакомо"
)

func main() {
	fmt.Println(know, i, your, name)
	fmt.Println(i, name, know, your)
	fmt.Println(your, know, name, i)
}
