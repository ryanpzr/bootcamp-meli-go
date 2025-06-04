package main

import "fmt"

func main() {
	numberMonth := 8

	mapMonth := map[int]string{
		1:  "Janeiro",
		2:  "Fevereiro",
		3:  "Março",
		4:  "Abril",
		5:  "Maio",
		6:  "Junho",
		7:  "Julho",
		8:  "Agosto",
		9:  "Setembro",
		10: "Outubro",
		11: "Novembro",
		12: "Dezembro",
	}

	for i, month := range mapMonth {
		if numberMonth == i {
			fmt.Printf("De acordo com o numero informado, o mês é: %s\n", month)
			return
		}
	}
}
