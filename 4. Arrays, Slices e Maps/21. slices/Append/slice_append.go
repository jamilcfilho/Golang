package main

import "fmt"

func main(){

	numeros := []int{2,4,6}
	fmt.Println(numeros)

	/* Comando <append> serve para adiconar elementos no slice.
	   Pode-se adicionar quantos elementos forém necessários
	   Possibilidade de *concatenar* slices dessa maneira também 
	*/
	numeros = append(numeros, 8,10,12)
	fmt.Println(numeros)
}