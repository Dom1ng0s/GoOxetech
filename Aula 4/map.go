package main

import "fmt"

func mapas() {
	idades := map[string]int{
		"Arthur": 24,
		"Davi":   19,
		"Pedro":  16,
	}
	fmt.Println(idades["Arthur"])
	fmt.Println(idades["Davi"])
	fmt.Println(idades["Pedro"])

	paises := map[string]string{
		"Arthur": "Alemanha",
		"Davi":   "Inglaterra",
		"Pedro":  "Grécia",
	}
	for pessoa, pais := range paises {
		fmt.Printf("O %s quer ir para %s\n", pessoa, pais)
	}
}
