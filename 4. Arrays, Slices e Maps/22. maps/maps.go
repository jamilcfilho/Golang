package main

import "fmt"

// Mapa = Uma coleção ordenada (lista) formada por pares (chave:valor)
// var x map[string]int     -> x é um mapa de strings para ints

func main() {

	mapa := make(map[string]int)
	mapa["chave1"] = 10
	mapa["chave2"] = 20
	mapa["chave3"] = 30


	// fmt.Println(mapa["chave2"])   -> Para exibir um valor específico do map

	fmt.Println(mapa) // Exibe uma lista do map, no caso a chave e o valor
}