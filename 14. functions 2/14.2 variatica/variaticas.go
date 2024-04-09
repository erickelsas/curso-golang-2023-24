package main

import "fmt"

func soma(numeros ...int) (total int) {
	total = 0
	for _, numero := range numeros {
		total += numero
	}

	return
}

func main() {
	fmt.Println(soma(1, 2, 3, 4, 5, 6, 7, 8))
}
