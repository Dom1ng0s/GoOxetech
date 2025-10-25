package main

import "fmt"

func exibirTipo(valor interface{}) {
	fmt.Printf("O valor '%v' é do tipo: %T\n", valor, valor)
}

func main() {
	exibirTipo(10)
	exibirTipo("Olá, mundo!")
	exibirTipo(true)
	exibirTipo(123.45)
	exibirTipo(nil)
	exibirTipo([]int{1, 2, 3})
}
