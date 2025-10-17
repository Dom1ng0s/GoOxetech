package main

import "fmt"

func main() {
	notas := []float64{7.0, 10.0, 9.0}

	if len(notas) == 0 {
		fmt.Println("Nenhuma nota fornecida")
		return
	}

	var soma float64
	for _, nota := range notas {
		soma += nota
	}

	media := soma / float64(len(notas))
	maior_nota := notas[0]
	menor_nota := notas[0]
	contador_media := 0

	fmt.Printf("Média: %.2f\n", media)
	for _, nota := range notas {
		if nota > media {
			contador_media++
		}
		if nota > maior_nota {
			maior_nota = nota
		}
		if nota < menor_nota {
			menor_nota = nota
		}
	}
	fmt.Printf("%d pessoas estão acima da média\n", contador_media)
	fmt.Printf("A maior nota foi %.2f\n", maior_nota)
	fmt.Printf("A menor nota foi %.2f\n", menor_nota)
}
