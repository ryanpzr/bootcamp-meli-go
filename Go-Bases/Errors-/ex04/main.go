package main

import (
	"fmt"
)

func checkSalary(salary int) error {
	var ErrorSalaryLess10000 = fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)

	if salary <= 100000 {
		return ErrorSalaryLess10000
	}
	return nil
}

func main() {
	salary := 2000

	if err := checkSalary(salary); err != nil {
		fmt.Println(err)
	} else if err == nil {
		fmt.Println("Error is Nil")
	} else {
		fmt.Println("Must pay tax")
	}
}
