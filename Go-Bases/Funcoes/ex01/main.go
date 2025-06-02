package main

import "fmt"

func main() {
	salario := 165.234
	result := getSalaryWithDiscount(float32(salario))
	fmt.Println(result)
}

func getSalaryWithDiscount(salary float32) float32 {
	if salary > 150.000 {
		salary = salary - (salary * 0.25)
		return salary
	} else if salary < 150.000 && salary > 50.000 {
		salary = salary - (salary * 0.10)
		return salary
	}

	return salary
}
