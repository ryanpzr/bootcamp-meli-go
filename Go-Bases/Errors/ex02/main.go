package main

import (
	"errors"
	"fmt"
)

var ErrorSalaryLess10000 = errors.New("Error: salary is less than 10000")

func checkSalary(salary int) error {

	if salary <= 100000 {
		return ErrorSalaryLess10000
	}
	return nil
}

func main() {
	salary := 2000

	if err := checkSalary(salary); errors.Is(err, ErrorSalaryLess10000) {
		fmt.Println(err)
	} else if err == nil {
		fmt.Println("Error is Nil")
	} else {
		fmt.Println("Must pay tax")
	}
}
