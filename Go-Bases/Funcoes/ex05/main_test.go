package main

import (
	"testing"
)

func TestDog(t *testing.T) {
	animalDog, msg := animal(dog)
	if msg != nil {
		return
	}
	var amount int
	amount += animalDog(10)
	expected := 100

	if amount != expected {
		t.Errorf("Era esperado %d e foi retornado %d", expected, amount)
	}
}

func TestCat(t *testing.T) {
	animalCat, msg := animal(dog)
	if msg != nil {
		return
	}
	var amount int
	amount += animalCat(5)

	expected := 50

	if amount != expected {
		t.Errorf("Era esperado %d e foi retornado %d", expected, amount)
	}
}
