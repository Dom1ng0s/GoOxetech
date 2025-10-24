package main

import "fmt"

type produto struct {
	nome  string
	preco float64
}

func (p *produto) AplicarDesconto(percentual float64) {
	p.preco -= p.preco * (percentual / 100)
}

func (p produto) Exibir() {
	fmt.Printf("Produto: %s\n\tPreço: %.2f\n", p.nome, p.preco)
}

func main() {
	fmt.Println("--- Teste 1: Livro sem desconto ---")
	livro := produto{nome: "A Arte da Guerra", preco: 50.00}

	livro.Exibir()

	fmt.Println("\n--- Teste 2: Laptop com desconto de 10% ---")
	laptop := produto{nome: "Laptop Gamer X500", preco: 4500.00}

	fmt.Printf("Preço original do %s: %.2f\n", laptop.nome, laptop.preco)

	laptop.AplicarDesconto(10.0)

	laptop.Exibir()

	fmt.Println("\n--- Teste 3: Mouse com dois descontos ---")
	mouse := produto{nome: "Mouse Sem Fio", preco: 100.00}

	mouse.Exibir()

	mouse.AplicarDesconto(20.0)
	fmt.Printf("Preço após primeiro desconto: %.2f\n", mouse.preco)

	mouse.AplicarDesconto(5.0)
	mouse.Exibir()
}
