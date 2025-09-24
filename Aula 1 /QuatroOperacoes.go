package main

import "fmt"

func main() {
	var a string = "Hello, World!"
	fmt.Println(a)
	var b, c int = 10, 20
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	f := "Curto"
	fmt.Println(f)

	println("A partir de agora começam as operações numéricas:")
	var g int = b + c
	fmt.Println(g)

	var h int = b - c
	fmt.Println(h)

	var i int = b * c
	fmt.Println(i)

	var j int = c / b
	fmt.Println(j)

}
