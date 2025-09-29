package main

import "fmt"

func main() {

	var a, b int
	var operador rune
	fmt.Scanf("%d %d", &a, &b)
	fmt.Scanf(" %c", &operador)

	switch operador {
	case '+':
		fmt.Printf("%d", a+b)
	case '-':
		fmt.Printf("%d", a-b)
	case '*':
		fmt.Printf("%d", a*b)
	case '/':
		if(b == 0){
			fmt.Print("Impossivel dividir por 0")
			return
		}
		var c float32 = float32(a) / float32(b)
		fmt.Printf("%f", c)
	default:
		fmt.Println("Operador inválido!")
	}


}
