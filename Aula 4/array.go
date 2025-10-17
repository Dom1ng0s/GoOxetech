package main

import "fmt"

func array() {
	var numeros [3]int
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	fmt.Println(numeros)
	// ARRAY
	nomes := [3]string{"Arthur", "Davi", "Pedro"}

	for i := 0; i < len(nomes); i++ {
		fmt.Println(i, nomes[i])
	}
	// OU:
	for f, nome := range nomes {
		fmt.Println(f, nome)
	}
}
