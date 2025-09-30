package main

import "fmt"

func main() {
	nomes := []string{"Arthur", "Davi", "Pedro"}
	for i, nomes := range nomes {
		fmt.Printf("Posição: %d - Nome:%s\n", i, nomes)
	}
}
