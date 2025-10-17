package main

import "fmt"

func map3() {
	notas := map[string]float64{
		"Arthur": 7.0,
		"Davi":   10.0,
		"Pedro":  9.0,
	}
	var soma float64
	for _, nota := range notas {
		soma += nota

	}
	media := float64(soma) / float64(len(notas))
	maior_nota := 0.00
	menor_nota := 10.00
	fmt.Println("Média:", media)
	for pessoa, nota := range notas {
		if nota > media {
			fmt.Printf("%s está na média\n", pessoa)
		}
		if nota > maior_nota {
			maior_nota = nota
		}
		if nota < menor_nota {
			menor_nota = nota
		}

	}
	fmt.Printf("A maior nota foi %f\n", maior_nota)
	fmt.Printf("A menor nota foi %f\n", menor_nota)

}
