package main

import "fmt"

type Retangulo struct {
	Comprimento int
	Largura     int
}

func (r Retangulo) Area() int {
	areaOperando := r.Comprimento * r.Largura
	fmt.Printf("A Área do Retângulo com dimensões: %d e %d é %d\n", r.Comprimento, r.Largura, areaOperando)
	return areaOperando
}

func (r Retangulo) DobrarArea() int {
	NovoOperando := Retangulo{Comprimento: r.Comprimento * 2, Largura: r.Largura * 2}
	fmt.Print("Novo Retângulo:")
	return NovoOperando.Area()
}

func AreaRetangulo() {
	ret1 := Retangulo{Comprimento: 10, Largura: 20}
	ret2 := Retangulo{Comprimento: 20, Largura: 30}
	ret3 := Retangulo{Comprimento: 30, Largura: 40}

	ret1.Area()
	ret2.Area()
	ret3.Area()

	ret1.DobrarArea()
	ret2.DobrarArea()
	ret3.DobrarArea()
}

func main() {
    AreaRetangulo()
}
