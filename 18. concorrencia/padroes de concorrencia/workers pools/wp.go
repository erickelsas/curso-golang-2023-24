package main

import "fmt"

func fibonacci(i int) int {
	if i <= 1 {
		return i
	}

	return fibonacci(i-2) + fibonacci(i-1)
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for tarefa := range tarefas {
		resultados <- fibonacci(tarefa)
	}
}

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}
