package main

import "fmt"

func dia(numero int) string {
	var diaDaSemana string
	switch numero {
	case 1:
		diaDaSemana = "Domingo"
	case 2:
		diaDaSemana = "Segunda-Feira"
	case 3:
		diaDaSemana = "Terça-Feira"
		fallthrough
	case 4:
		diaDaSemana = "Quarta-Feira"
	case 5:
		diaDaSemana = "Quinta-Feira"
	case 6:
		diaDaSemana = "Sexta-Feira"
	case 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Inválido"
	}

	return diaDaSemana
}

func dia2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Inválido"
	}
}

func main() {
	numero := 3

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if n := numero; n > 0 {
		fmt.Println(numero)
		fmt.Println(dia(n))
	}
}
