package main

import (
	"aula5/models"
	"fmt"
)

func Livraria() {
	livros := []models.Livro{
		{Titulo: "Senhor dos Aneis", Autor: "J.R.R Tolkien", Ano: 1954},
		{Titulo: "Harry Potter", Autor: "J.K. Rowling", Ano: 1997},
	}
	var titulo, autor string
	var ano int
	fmt.Println("Os Livros que temos Atualmente: \n", livros)
	fmt.Printf("Crie um Livro\n Titulo: ")
	fmt.Scanf("%s\n", &titulo)
	fmt.Printf("Autor: ")
	fmt.Scanf("%s\n", &autor)
	fmt.Printf("Ano:")
	fmt.Scanf("%d\n", &ano)
	livros = append(livros, models.Livro{Titulo: titulo, Autor: autor, Ano: ano})
	fmt.Println(livros)
	livroAtualizado := livros[1]
	livroAtualizado.Titulo = "Harry Potter e a camara secreta"
	livroAtualizado.Ano = 1997
	livros = append(livros, livroAtualizado)
	fmt.Println(livros)
}
