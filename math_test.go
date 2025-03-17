package main

import "testing"

func TestSoma(t *testing.T) {

	result := Soma(15, 5)

	if result != 20 {
		t.Errorf("Resultado da soma é invalido: Resultado %d, Esperado: %d", result, 20)
	}
	
}