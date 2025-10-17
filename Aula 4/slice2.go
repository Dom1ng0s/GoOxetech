package main

import "fmt"

func slice2() {
	slice := []string{"A", "B", "C", "D", "E"}

	for i, j := range slice {
		fmt.Printf("Indice: %d Nome %s\n", i, j)
	}
	slice = append(slice, "F", "G")
	for i, j := range slice {
		fmt.Printf("Indice: %d Nome %s\n", i, j)
	}
}
