package main

import (
	"errors"
	"fmt"
)

func main() {
	var n int = 100
	fmt.Println(n)

	//int sem sinal (só positivos)
	var n2 uint = 10000
	fmt.Println(n2)

	//alias
	//int32 = rune
	var n3 rune = 12345
	fmt.Println(n3)

	//uint8 = byte
	var n4 byte = 123
	fmt.Println(n4)

	var nr float32 = 123.45
	fmt.Println(nr)

	var nr2 float64 = 12300000.45
	fmt.Println(nr2)

	nr3 := 1234.65
	fmt.Println(nr3)

	var str string = "Texto"
	fmt.Println(str)

	char := '%'
	fmt.Println(char)

	var texto int
	fmt.Println(texto)

	var bool1 bool
	fmt.Println(bool1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
