package main

import "fmt"

func ivtSinal(n int) int {
	return n * -1
}

func ivtSinalPont(n *int) {
	*n = *n * -1
}

func main() {
	numero := 20
	nI := ivtSinal(numero)

	fmt.Println(nI)
	fmt.Println(numero)

	fmt.Println(numero)
	ivtSinalPont(&numero)
	fmt.Println(numero)
}
