package main

import "fmt"

type Error struct {
	message    string
	statusCode string
}

func main() {
	salary, err := calculateSalary(160, 470)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("O salário final é: R$%d\n", salary)
}

func (e *Error) myError() error {
	return fmt.Errorf("Message: %s | StatusCode: %s", e.message, e.statusCode)
}

func calculateSalary(hoursWorked int, hoursWorkedPrice int) (int, error) {
	totalSalary := hoursWorked * hoursWorkedPrice

	if totalSalary >= 150000 {
		totalSalary = totalSalary + (totalSalary * 10 / 100)
		return totalSalary, nil
	} else if hoursWorked < 80 || hoursWorked < 0 {
		e := Error{message: "the worker cannot have worked less than 80 hours per month", statusCode: "404"}
		return 0, e.myError()
	}

	return totalSalary, nil
}
