package main

import (
	"fmt"
	enderecos "intro-testes/endereco"
)

func main() {
	tipoEndereco := enderecos.Tipo("Avenida Paulista")
	fmt.Println(tipoEndereco)
}
