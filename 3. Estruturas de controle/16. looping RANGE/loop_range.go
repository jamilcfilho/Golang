package main

import "fmt"

// Comando muito utilizado para PERCORRER arrays, slices, maps e strings
func main(){

	nomes := []string{"Ana", "Bruno", "Carlos"}
	for indice, nome := range nomes{
		fmt.Println(indice, nome)
	}

	// Quando não precisa utilizar o Índice (valor na frente do nome) pode-se utilizar o <_>
	for _, nome := range nomes{
		fmt.Println(nome)
	}
}