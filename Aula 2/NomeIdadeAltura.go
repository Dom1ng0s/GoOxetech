package main

import "fmt"

const Pi = 3.14159

func main() {
	var nome string
	var idade int8
	var altura uint8
	fmt.Println("Insira seu nome, altura(Centimetros) e idade")
	fmt.Scanf("%s\n", &nome)
	fmt.Scanf("%d\n", &altura)
	fmt.Scanf("%d\n", &idade)

	fmt.Printf("Seu nome é %s, você tem %d anos e sua altura é %d centimetros\n", nome, idade, altura)
	if idade < 18 {
		fmt.Println("Voce e de menor")
	} else {
		fmt.Println("Voce e de maior")
	}
}
