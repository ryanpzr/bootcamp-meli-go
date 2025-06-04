package main

import "fmt"

type Cliente struct {
	nome            string
	idade           int
	tempo_empregado float64
	salario         float64
}

func main() {
	cliente01 := Cliente{
		nome:            "Ryan",
		idade:           25,
		tempo_empregado: 6,
		salario:         1426.7,
	}

	cliente02 := Cliente{
		nome:            "Marcos",
		idade:           35,
		tempo_empregado: 3,
		salario:         1926,
	}

	cliente03 := Cliente{
		nome:            "Henrique",
		idade:           29,
		tempo_empregado: 0.5,
		salario:         3450,
	}

	var listaClientes []Cliente
	listaClientes = append(listaClientes, cliente01)
	listaClientes = append(listaClientes, cliente02)
	listaClientes = append(listaClientes, cliente03)

	for _, cliente := range listaClientes {
		validaEmprestimo(cliente)
	}
}

func validaEmprestimo(c Cliente) {
	if c.idade > 22 && c.tempo_empregado > 1 {
		if c.salario > 1000 {
			fmt.Printf("Foi liberado empréstimo ao cliente %s, com juros\n", c.nome)
			return
		}
		fmt.Printf("Foi liberado empréstimo ao cliente %s, sem juros\n", c.nome)
		return
	}

	fmt.Printf("Não foi liberado empréstimo ao cliente %s, pois não atende aos requisítos\n", c.nome)
}
