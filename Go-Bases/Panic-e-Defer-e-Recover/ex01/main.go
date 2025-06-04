package main

import (
	"fmt"
	"os"
)

func main() {
	txt, err := os.ReadFile("customers.txt")
	defer fmt.Println("execução concluida")
	if err != nil {
		panic("he indicated file was not found or is damaged")
	}

	fmt.Println(txt)
}
