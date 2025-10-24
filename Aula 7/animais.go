package main

import "fmt"

type Animal interface {
	Falar() string
}

func EmitirSom(a Animal) {
	fmt.Println(a.Falar())
}

type Cachorro struct{}
type Gato struct{}
type Vaca struct{}

func (c Cachorro) Falar() string {
	return "Au au!"
}
func (g Gato) Falar() string {
	return "Miau!"
}
func (v Vaca) Falar() string {
	return "muuu!"
}

func animais() {
	c := Cachorro{}
	g := Gato{}
	v := Vaca{}

	EmitirSom(c)
	EmitirSom(g)
	EmitirSom(v)
}
