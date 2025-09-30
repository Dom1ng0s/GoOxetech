package main

import "fmt"

func main() {
	resultado := somar(10, 1)
	fmt.Println(resultado)
}

func somar(a int, b int) int {
	return a + b
}
