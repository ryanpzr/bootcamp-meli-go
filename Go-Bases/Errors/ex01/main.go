package main

import (
	"errors"
	"fmt"
)

func checkSalary(salary int) error {
	if salary < 150000 {
		return errors.New("Error: the salary entered does not reach the taxable minimum")
	}
	return nil
}

func main() {
	salary := 2000

	if err := checkSalary(salary); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Must pay tax")
	}
}
