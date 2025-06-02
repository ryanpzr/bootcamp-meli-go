package main

import (
	"errors"
	"fmt"
)

const (
	dog = "dog"
	cat = "cat"
)

func main() {
	animalDog, msg := animal(dog)
	if msg != nil {
		return
	}
	animalCat, msg := animal(cat)
	if msg != nil {
		return
	}

	var amount int

	amount += animalDog(10)
	amount += animalCat(10)

	fmt.Printf("São necessarios %dkg de comida para os animais.", amount)
}

func animal(animal string) (func(int) int, error) {
	switch animal {
	case "dog":
		return func(quantidadeAnimal int) int {
			return quantidadeAnimal * 10
		}, nil
	case "cat":
		return func(quantidadeAnimal int) int {
			return quantidadeAnimal * 5
		}, nil
	default:
		return nil, errors.New("Animal não encontrado.")
	}
}
