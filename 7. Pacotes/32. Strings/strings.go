// Pacote que serve para manipulação de palavras/strings

package main

import (
	"fmt"
	"strings"
)

func main() {
	texto := "Golang é uma linguagem incrível e simples"

	fmt.Println(strings.ToUpper(texto)) // Tranforma o texto tudo em MAIÚSCULA
	fmt.Println(strings.ToLower(texto)) // Transforma o texto tudo em minúscula
	fmt.Println(strings.Contains(texto, "simples")) // Verifica se um texto/trecho/pedaço contém dentro da frase toda e retorna um booleano
	fmt.Println(strings.ReplaceAll(texto, "Golang", "Go")) // Substitui trechos da frase

	palavras := strings.Split(texto, " ") // Serve para separar palavras
	fmt.Println("Palavras:", palavras)

	novoTexto := strings.Join(palavras, "-") // Serve para unir palavras
	fmt.Println("Novo texto:", novoTexto)
}