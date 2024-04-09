package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circle struct {
	raio float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circle{10}
	escreverArea(c)
}
