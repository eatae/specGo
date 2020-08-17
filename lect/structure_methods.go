package main

import "fmt"

/* SomeStudent */
type SomeStudent struct {
	name string
	year int
}

/* method */
func (s SomeStudent) printInfo() {
	fmt.Println("Name: ", s.name)
	fmt.Println("Year: ", s.year)
}

/* value receiver */
func (s SomeStudent) valueReceive() { // внутри метода работаем с копией SomeStudent
	s.year = 1000
	fmt.Println("Into value method: ", s.year)
}

/* pointer receiver */
func (s *SomeStudent) pointerReceive() { // внутри метода работаем с экземпляром SomeStudent
	s.year = 2000
	fmt.Println("Into pointer method: ", s.year)
}

func main() {
	student := SomeStudent{
		name: "John",
		year: 35,
	}
	student.printInfo()
	/****************/
	student.valueReceive()
	fmt.Println("After value method: ", student.year) // значение экземпляра не изменилось
	student.pointerReceive()
	fmt.Println("After pointer method: ", student.year) // значение экземпляра изменилось
}
