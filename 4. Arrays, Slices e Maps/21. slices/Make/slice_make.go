package main

import "fmt"

func main(){

	fatia1 := make([]int, 3) // Tamanho do slice é 3 com capacidade de 3
	fmt.Println(fatia1)

	fatia2 := make([]int, 4, 6) // Tamanho do slice é 4 com capacidade de 6
	fmt.Println(fatia2)
}