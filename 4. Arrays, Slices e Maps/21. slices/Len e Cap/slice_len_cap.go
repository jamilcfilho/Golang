package main

import "fmt"

func main(){

	fatia1 := make([]int, 3)
	fmt.Println(fatia1)

	fatia2 := make([]int, 4, 6)
	fmt.Println(fatia2)

	// É possível descobrir o tamanho e a capacidade do slice, utilizando os comandos <len> e <cap>
	fmt.Println(len(fatia1))
	fmt.Println(cap(fatia1))

	fmt.Println(len(fatia2))
	fmt.Println(cap(fatia2))
}