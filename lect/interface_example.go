package main

import "fmt"

type EmptyInterface interface {
}

type SalaryCalculator interface {
	CalcSalary() int
}

/* Permanent постоянный сотрудник */
type Permanent struct {
	salary    int
	premiales int
}

func (p Permanent) CalcSalary() int {
	return p.salary + p.premiales
}

/* Contract контрактный сотрудник */
type Contract struct {
	fixSalary int
}

func (c Contract) CalcSalary() int {
	return c.fixSalary
}

/* Freelancer */
type Freelancer struct {
	perHour int
	hours   int
}

func (f Freelancer) CalcSalary() int {
	return f.perHour * f.hours
}

func TotalSalary(workers []SalaryCalculator) int {
	var total int
	for _, worker := range workers {
		total += worker.CalcSalary()
	}
	return total
}

func AnyArgument(i interface{}) {
	fmt.Printf("Item type: %T, value: %v\n", i, i)
}

func AnyArgument2(i interface{}) {
	intValue, ok := i.(int)
	if !ok {
		fmt.Print("Value is not integer", intValue)
	}
	fmt.Println(intValue)
}

func main() {

	i := 123
	s := "some string"

	AnyArgument(i)
	AnyArgument(s)

	var salary SalaryCalculator
	fmt.Printf("%T\n", salary)

	AnyArgument2("fooBar")

	permanent := Permanent{120, 20}
	contract := Contract{130}
	freelancer := Freelancer{20, 7}

	fmt.Println(TotalSalary([]SalaryCalculator{permanent, contract, freelancer}))

	fmt.Printf("Item type: %T, value: %v\n", freelancer, freelancer)
}
