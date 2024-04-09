package main

import "fmt"

func closure() func() {
	txt := "Dentro da Closure"

	funcao := func() {
		fmt.Println(txt)
	}

	return funcao
}

func main() {
	txt := "Dentro da Main"
	fmt.Println(txt)

	f := closure()
	f()
}
