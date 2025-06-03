package main

import "fmt"

func main() {
	minutosTrabalhados := 3600
	result := getSalaryFromCategory(minutosTrabalhados, "B")
	fmt.Printf("O salario final é: US$%f\n", result)
}

func getSalaryFromCategory(minutes int, category string) float64 {
	salary := 0.0
	additional := 0.0
	switch category {
	case "B":
		hoursWorked := transformMinutesInHours(minutes)
		salary = 1.500
		additional = 0.20
		salary = (salary * float64(hoursWorked))
		salary = salary + (salary * additional)
	case "C":
		hoursWorked := transformMinutesInHours(minutes)
		salary = 1.000
		salary = (salary * float64(hoursWorked))
	case "A":
		hoursWorked := transformMinutesInHours(minutes)
		salary = 3.000
		additional = 0.50
		salary = (salary * float64(hoursWorked))
		salary = salary + (salary * additional)
	}
	return salary
}

func transformMinutesInHours(minutes int) int {
	return minutes / 60
}
