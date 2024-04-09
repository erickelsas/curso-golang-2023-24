package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int16
}

func main() {
	var u usuario
	u.nome = "Erick"
	u.idade = 20
	fmt.Println(u)

	enderecoEx := endereco{"Rua dos Bobos", 0}

	u1 := usuario{"Davi", 21, enderecoEx}
	fmt.Println(u1)

	u2 := usuario{nome: "João"}
	fmt.Println(u2)
}
