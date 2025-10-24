package main

import "fmt"

type Retangulo struct {
	Comprimento int
	Largura     int
}

func Area(Operando Retangulo) int {
	areaOperando := Operando.Comprimento * Operando.Largura
	fmt.Printf("A Area do Retângulo com dimensões: %d e %d é %d\n", Operando.Comprimento, Operando.Largura, areaOperando)
	return areaOperando
}
func DobrarArea(Operando Retangulo) int {
	NovoOperando := Retangulo{Comprimento: Operando.Comprimento * 2, Largura: Operando.Largura * 2}
	fmt.Print("Novo Retângulo:")
	return Area(NovoOperando)

}

func main() {
	ret1 := Retangulo{Comprimento: 10, Largura: 20}
	ret2 := Retangulo{Comprimento: 20, Largura: 30}
	ret3 := Retangulo{Comprimento: 30, Largura: 40}
	Area(ret1)
	Area(ret2)
	Area(ret3)
	DobrarArea(ret1)
	DobrarArea(ret2)
	DobrarArea(ret3)
}
