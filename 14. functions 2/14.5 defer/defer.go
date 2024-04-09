package main

import "fmt"

func f1() {
	fmt.Println("Exemplo f1")
}

func f2() {
	fmt.Println("Exemplo f2")
}

func aprovado(n1, n2 float32) bool {
	media := (n1 + n2) / 2

	defer fmt.Println("Média calculada. Resultado será retornado!")
	if media < 6 {
		return false
	}

	return true
}

func main() {
	fmt.Println(aprovado(5, 8))
}
