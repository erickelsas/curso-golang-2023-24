package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array [5]int
	fmt.Println(array)

	array1 := [5]int{0, 2, 4, 6, 8}
	fmt.Println(array1)

	array2 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array2)

	slice := []int{0, 1, 2}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array))

	slice = append(slice, 3)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = 20
	fmt.Println(slice2)

	//arrays internos
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
