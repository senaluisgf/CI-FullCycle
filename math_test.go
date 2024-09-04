package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado é inválido: Resultado %d. Esperado: %d.", total, 30)
	}
}

func TestSub(t *testing.T) {
	total := sub(15, 15)

	if total != 0 {
		t.Errorf("Resultado é inválido: Resultado %d. Esperado: %d.", total, 0)
	}
}

func TestTimes(t *testing.T) {
	total := times(15, 15)

	if total != 225 {
		t.Errorf("Resultado é inválido: Resultado %d. Esperado: %d.", total, 225)
	}
}

func TestDiv(t *testing.T) {
	total := div(15, 15)

	if total != 1 {
		t.Errorf("Resultado é inválido: Resultado %d. Esperado: %d.", total, 1)
	}
}

func TestPow(t *testing.T) {
	total := pow(2, 3)

	if total != 8 {
		t.Errorf("Resultado é inválido: Resultado %d. Esperado: %d.", total, 8)
	}
}
