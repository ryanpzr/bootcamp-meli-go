package main

import (
	"testing"
)

func TestGetSalaryWithDiscountLessThan50Thousand(t *testing.T) {
	salario01 := 40.000

	result01 := getSalaryWithDiscount(float32(salario01))

	salarioEsperado01 := 40.000

	if result01 != float32(salarioEsperado01) {
		t.Errorf("Não foi aplicado corretamente o disconto do imposto sobre o salário de R$%f", salario01)
	}
}

func TestGetSalaryWithDiscountGreaterThan50Thousand(t *testing.T) {
	salario02 := 60.000

	result02 := getSalaryWithDiscount(float32(salario02))

	salarioEsperado02 := 54.000

	if result02 != float32(salarioEsperado02) {
		t.Errorf("Não foi aplicado corretamente o disconto do imposto sobre o salário de R$%f", salario02)
	}
}

func TestGetSalaryWithDiscountGreaterThan150Thousand(t *testing.T) {
	salario03 := 180.000

	result03 := getSalaryWithDiscount(float32(salario03))

	salarioEsperado03 := 135.000

	if result03 != float32(salarioEsperado03) {
		t.Errorf("Não foi aplicado corretamente o disconto do imposto sobre o salário de R$%f", salario03)
	}
}
