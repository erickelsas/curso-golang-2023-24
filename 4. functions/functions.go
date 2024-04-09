package main

import "fmt"

func somar(x int, y int) int {
	return x + y
}

func calculosMat(x, y int) (int, int) {
	return x + y, x - y
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Função 1")

	soma, sub := calculosMat(3, 2)
	fmt.Println(soma, sub)

	soma1, _ := calculosMat(4, 5)
	fmt.Println(soma1)
}
