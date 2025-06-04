package main

import (
	"errors"
	"fmt"
)

type Cliente struct {
	File        string
	Name        string
	Id          int
	PhoneNumber string
	Address     string
}

var clientList []Cliente

func main() {
	defer fmt.Println("End of execution")
	defer fmt.Println("Several errors were detected at runtime")

	mountClientList()

	newClient := Cliente{
		File:        "foto",
		Name:        "Ryan",
		Id:          3,
		PhoneNumber: "48996480085",
		Address:     "Rua Professora Ryan Campos, 393",
	}

	newClient02 := Cliente{
		File:        "foto",
		Name:        "Julio",
		Id:          4,
		PhoneNumber: "48997673277",
		Address:     "Rua Professora Julio Santos, 393",
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado do panic:", r)
		}
	}()

	cliente, err := validateClientInfo(newClient)
	if err != nil {
		fmt.Println(err)
	} else {
		isRegister(cliente)
		clientList = append(clientList, cliente)
	}

	cliente02, err := validateClientInfo(newClient02)
	if err != nil {
		fmt.Println(err)
	} else {
		isRegister(cliente02)
		clientList = append(clientList, cliente02)
	}

	for i := range clientList {
		fmt.Printf("%+v\n", clientList[i])
	}
}

func mountClientList() {
	cliente1 := Cliente{
		File:        "foto",
		Name:        "Ryan",
		Id:          1,
		PhoneNumber: "48996480085",
		Address:     "Rua Professora Rosinha Campos, 393",
	}

	cliente2 := Cliente{
		File:        "foto",
		Name:        "Marcos",
		Id:          2,
		PhoneNumber: "48997864456",
		Address:     "Rua Professor Bayer Filho, 38",
	}

	clientList = append(clientList, cliente1)
	clientList = append(clientList, cliente2)
}

func isRegister(newClient Cliente) bool {
	for _, cliente := range clientList {
		if cliente.Name == newClient.Name && cliente.PhoneNumber == newClient.PhoneNumber {
			panic("client already exists")
		}
	}

	return false
}

func validateClientInfo(cl Cliente) (Cliente, error) {
	if cl.Id == 0 {
		return Cliente{}, errors.New("O campo Id não pode ser 0")
	} else if cl.PhoneNumber == "0" {
		return Cliente{}, errors.New("O campo Phone não pode ser 0")
	}

	return cl, nil
}
