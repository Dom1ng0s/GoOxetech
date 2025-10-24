package main

import "fmt"

type Desenhavel interface {
	Desenhar()
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Desenhar() {
	fmt.Println("Desenhando Circulo com o Raio", c.Raio)
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (r Retangulo) Desenhar() {
	fmt.Println("Desenhando Retangulo com medidas", r.Altura, r.Largura)
}

func MostrarDesenho(d Desenhavel) {
	d.Desenhar()
}

func main() {
	circulo := Circulo{Raio: 5}
	retangulo := Retangulo{Altura: 5, Largura: 10}
	MostrarDesenho(circulo)
	MostrarDesenho(retangulo)

}
