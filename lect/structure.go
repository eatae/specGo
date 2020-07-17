package main

import "fmt"



func main() {
	initStruct()
}


/* Employee structure */
type Employee struct {
	firstName string
	lastName  string
	age       int
}

/* Professor structure - fields as type */
type Professor struct {
	string
	int
}


/* Adress struc (for nested) */
type Address struct {
	city   string
	street string
}

/* Person struct */
type Person struct {
	name string
	age  int
	// Nested structure
	address Address
}



func initStruct() {

	/* create */
	empl1 := Employee{
		firstName: "Bob",
		age:       25,
		lastName:  "Dub",
	}


	/* create */
	empl2 := Employee{"Alice", "Yandex", 20}


	/* anonymous structure */
	empl3 := struct {
		firstName string
		lastName  string
	}{
		firstName: "Georg",
		lastName:  "Tomp",
	}


	/* empty structure */
	var empl4 Employee
	empl4.firstName = "John"


	/* structure pointer */
	emplPointer := &Employee{"Pointer", "Snow", 34}
	emplPointer.age = 45 // set
	/* structure no pointer */
	emplNoPointer := *emplPointer
	emplNoPointer.firstName = "NoPointer"


	/* nested structure */
	person := Person{
		name: "Bob",
		age:  25,
		address: Address{
			city:   "Moscow",
			street: "Ladozhskaya",
		},
	}



	fmt.Println("____________INIT____________")
	fmt.Println(empl1, empl2, empl3)
	fmt.Println(empl4)
	fmt.Println(empl4.age)
	fmt.Println("____________POINTER____________")
	fmt.Println(emplPointer)   // Pointer
	fmt.Println(*emplPointer)  // Struct
	fmt.Println(emplNoPointer) // Struct No Pointer
	fmt.Println("____________NESTED_STRUCTURE____________")
	fmt.Println("Person", person)
	fmt.Println("Person.name", person.name)
	fmt.Println("Person.Address", person.address)
	fmt.Println("Person.Address.street", person.address.street)
}

