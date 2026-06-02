package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(7, 7)

	if total != 14 {
		t.Errorf("Resultado da conta é inválido: Resultado %d. Esperado %d", total, 14)
	}
}