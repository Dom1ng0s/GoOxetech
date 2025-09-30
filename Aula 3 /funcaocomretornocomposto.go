package main

import "fmt"

func main() {
	valido, resultado := divisao(10, 0)
	fmt.Printf("O calculo é válido? %v, o resultado é: %.2f", valido, resultado)
}

func divisao(a float64, b float64) (bool, float64) {
	if b == 0 {
		return false, 0
	}
	return true, a / b
}
