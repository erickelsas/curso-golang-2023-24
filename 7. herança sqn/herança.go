package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint16
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p := pessoa{"Erick", "Freitas", 20, 184}
	e := estudante{p, "Eng de Software", "UTFPR"}

	fmt.Println(e)
}
