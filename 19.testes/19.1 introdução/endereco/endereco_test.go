package enderecos

import "testing"

type teste struct {
	inserido string
	esperado string
}

func TestTipo(t *testing.T) {
	cenarios := []teste{
		{"Rua ABC", "Rua"},
		{"Avenida", "Avenida"},
		{"Rodovia BR369", "Rodovia"},
		{"Estrada João Melão", "Estrada"},
		{"Praça das Rosas", "Inválido"},
	}

	for _, cenario := range cenarios {
		tipoRecebido := Tipo(cenario.inserido)

		if tipoRecebido != cenario.esperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoRecebido, cenario.esperado)
		}
	}
}
