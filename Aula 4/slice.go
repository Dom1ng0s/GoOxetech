package main

import "fmt"

func slice() {
	slice := []int{1, 2}
	slice = append(slice, 3, 4) // retorna o slice com o valor adicionado
	fmt.Println(slice, len(slice))
}
