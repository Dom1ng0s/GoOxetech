package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Scanf("%d %d", &n1, &n2)
	resultado := maior(n1, n2)
	fmt.Println(resultado)

}

func maior(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
