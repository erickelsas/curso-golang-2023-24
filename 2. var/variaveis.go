package main

import "fmt"

func main() {
	var variavel1 string = "Var 1"
	fmt.Println(variavel1)

	variavel2 := "Var 2" //inferencia de tipo
	fmt.Println(variavel2)

	var (
		variavel3 string = "Var 3"
		variavel4 string = "Var 4"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Var 5", "Var 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
