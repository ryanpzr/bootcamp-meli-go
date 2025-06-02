package main

import "fmt"

func main() {
	minutosTrabalhados := 3600
	result := getSalaryFromCategory(minutosTrabalhados, "B")
	fmt.Printf("O salario final é: US$%d\n", result)
}

func getSalaryFromCategory(minutes int, category string) int {
	salary := 0
	switch category {
	case "B":
		hoursWorked := transformMinutesInHours(minutes)
		salary = 1500
		additional := 0.20
		salary = (salary * hoursWorked)
		salary = salary + (salary * int(additional))
	case "C":
		hoursWorked := transformMinutesInHours(minutes)
		salary = 1000
		salary = (salary * hoursWorked)
	case "A":
		hoursWorked := transformMinutesInHours(minutes)
		salary = 3000
		additional := 0.50
		salary = (salary * hoursWorked)
		salary = salary + (salary * int(additional))
	}
	return salary
}

func transformMinutesInHours(minutes int) int {
	return minutes / 60
}
