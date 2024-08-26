package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.ErrorF("Resultado é inválido: Resultado %d. Esperado: %d.", total, 30)
	}
}
