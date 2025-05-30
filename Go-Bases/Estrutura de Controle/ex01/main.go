package main

import "fmt"

func main() {
	palavra := "Ryan"
	letras := []string{}
	fmt.Println(len(palavra))
	for _, pl := range palavra {
		letras = append(letras, string(pl))
	}

	fmt.Println(letras)
}
