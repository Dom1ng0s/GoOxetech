package main

import "fmt"

func main() {
	palavra := "Dom1ng0s"
	for posicao, letra := range palavra {
		fmt.Printf("Posição: %d Letra: %c\n", posicao, letra)
	}
}
