package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	qntFuncionarioMaiorQue21 := 0
	employees["Frederico"] = 25

	for name, idade := range employees {
		switch {
		case name == "Benjamin":
			fmt.Printf("Benjamin tem %d anos.\n", idade)
		case idade > 21:
			qntFuncionarioMaiorQue21++
		case name == "Pedro":
			delete(employees, name)
		default:
			continue
		}
	}

	fmt.Printf("Possui %d funcionarios maiores que 21 anos.\n", qntFuncionarioMaiorQue21)
}
