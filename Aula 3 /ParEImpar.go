package main

import "fmt"

func main() {
	var numero int64
	fmt.Scanf("%d", &numero)
	parEImpar(int(numero))
}

func parEImpar(a int) {
	for i := 1; i < a; i++ {
		if i%2 == 0 {
			fmt.Printf("%d é par\n", i)
		} else {
			fmt.Printf("%d é impar\n", i)
		}

	}
}
