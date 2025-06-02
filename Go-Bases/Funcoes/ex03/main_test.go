package main

import "testing"

func TestGetSalaryFromCategoryA(t *testing.T) {
	resultCategoryA := getSalaryFromCategory(3600, "A")

	expectedA := 270.000000

	if resultCategoryA != expectedA {
		t.Errorf("Era esperado %f e foi retornado %f", expectedA, resultCategoryA)
	}
}

func TestGetSalaryFromCategoryB(t *testing.T) {
	resultCategoryB := getSalaryFromCategory(3600, "B")

	expectedB := 108.000000

	if resultCategoryB != expectedB {
		t.Errorf("Era esperado %f e foi retornado %f", expectedB, resultCategoryB)
	}
}

func TestGetSalaryFromCategoryC(t *testing.T) {
	resultCategoryC := getSalaryFromCategory(3600, "C")

	expectedC := 60.000000

	if resultCategoryC != expectedC {
		t.Errorf("Era esperado %f e foi retornado %f", expectedC, resultCategoryC)
	}
}
