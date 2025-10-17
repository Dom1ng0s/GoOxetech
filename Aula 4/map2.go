package main

import "fmt"

func map2() {
	idades := map[string]int{
		"Arthur": 24,
		"Davi":   19,
		"Pedro":  16,
	}
	idades["José"] = 31

	idade, existe := idades["Davi"]

	if existe {
		fmt.Printf("Davi tem %d anos\n", idade)
	} else {
		fmt.Printf("Davi Não existe")
	}
	delete(idades, "José")
	fmt.Println(idades)
}
