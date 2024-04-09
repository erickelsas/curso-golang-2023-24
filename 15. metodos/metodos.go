package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando dados de %s\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) aniversario() {
	u.idade++
}

func main() {
	u := usuario{"Usuário 1", 20}
	fmt.Println(u)
	u.salvar()
	fmt.Println(u.maiorDeIdade())

	u.aniversario()
	fmt.Println(u.idade)
}
