package main

import "fmt"

func main() {
	valido, resultado, operacao := divisao(10, 0)
	fmt.Printf("A operação é: %s, O calculo é válido? %v, o resultado é: %.2f", operacao, valido, resultado)
}

func divisao(a float64, b float64) (bool, float64, string) {
	if b == 0 {
		return false, 0, "divisão"
	}
	return true, a / b, "divisão"
}
