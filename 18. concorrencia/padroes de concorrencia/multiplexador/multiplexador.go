package main

import (
	"fmt"
	"math/rand"
	"time"
)

func escrever(txt string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", txt)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}

func multiplexar(canal1, canal2 <-chan string) <-chan string {
	canalS := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canal1:
				canalS <- mensagem
			case mensagem := <-canal2:
				canalS <- mensagem
			}
		}
	}()

	return canalS
}

func main() {
	canal := multiplexar(escrever("Olá mundo"), escrever("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}
