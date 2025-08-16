package main

import "fmt"

// escopo, que pode ser utilizado depois em qualquer lugar do arquivo
var saudacao string = "Hello world!"

// bloco, na qual se encontra a função onde é a unidade fundamental de execução qu representa os comandos do programa
func main() {
	
	fmt.Println(saudacao)
	
}