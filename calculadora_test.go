package main

import "testing"

func TestSomar(t *testing.T) {
	total := somar(10, 10)
	if total != 20 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: 30", total)
	}
}
