package main

import "testing"

func TestSomar(t *testing.T) {
	total := somar(15, 15)
	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: 30", total)
	}
}
