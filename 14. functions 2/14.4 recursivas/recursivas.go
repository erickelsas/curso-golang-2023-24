package main

import "fmt"

func fibonacci(i uint) uint {
	if i <= 1 {
		return i
	}

	return fibonacci(i-2) + fibonacci(i-1)
}

func main() {
	posicao := uint(10)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
