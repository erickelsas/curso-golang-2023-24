package main

import "fmt"

func main() {
	var v1 int = 10
	var v2 int = v1
	fmt.Println(v1, v2)

	v1++
	fmt.Println(v1, v2)

	var p *int = &v1
	fmt.Println(*p)

	*p++
	fmt.Println(v1)

}
