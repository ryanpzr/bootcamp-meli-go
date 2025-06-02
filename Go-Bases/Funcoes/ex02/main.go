package main

import "fmt"

func main() {
	result := getMediaNotas(9, 9, 9, 2, 5, 7, 7, 8, 9, 10, 4)
	fmt.Printf("A média da turma é: %d\n", result)
}

func getMediaNotas(notas ...int) int {
	somaNotas := 0
	for _, nota := range notas {
		somaNotas += nota
	}

	return somaNotas / len(notas)
}
