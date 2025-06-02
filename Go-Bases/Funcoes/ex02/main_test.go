package main

import "testing"

func TestMediaNotas(t *testing.T) {
	result := getMediaNotas(7, 5, 9, 2, 6, 4, 7)

	expected := 5

	if result != expected {
		t.Errorf("A média da turma não foi calculada corretamente, deveria ser %d e foi retornado %d", expected, result)
	}
}
