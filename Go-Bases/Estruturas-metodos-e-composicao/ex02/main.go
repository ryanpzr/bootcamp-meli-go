package main

import (
	"fmt"
)

type Person struct {
	id          int
	name        string
	dateOfBirth string
}

type Employee struct {
	id       int
	position int
	Person
}

func (e Employee) PrintEmployee() {
	fmt.Println(e.Person.name, e.Person.dateOfBirth)
}

func main() {
	person := Person{
		id:          1,
		name:        "Ryan",
		dateOfBirth: "07/10/2004",
	}

	employee := Employee{
		id:       1,
		position: 1,
		Person:   person,
	}

	employee.PrintEmployee()
}
