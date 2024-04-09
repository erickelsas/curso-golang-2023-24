package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
		time.Sleep(time.Second)
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for _, letra := range "PALAVRA" {
		fmt.Println(string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
