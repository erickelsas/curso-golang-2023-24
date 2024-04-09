package enderecos

import "strings"

func Tipo(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	primeiraPalavra := strings.ToLower(strings.Split(endereco, " ")[0])

	tipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			tipoValido = true
		}
	}

	if tipoValido {
		return strings.Title(primeiraPalavra)
	}

	return "Inválido"
}
