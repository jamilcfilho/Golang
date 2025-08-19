package main

import "fmt"

// Estruturas = São coleções de "Campos nomeados" que agrupa valores para formar entidades (Pessoas, Produtos, Endereço, etc.) e organiza os dados de forma clara

// type <nome da estrutura> struct   -> Sintaxe padrão

type pessoa struct {
	nome string
	idade int
	profissao string
}

func main() {

	p1 := []pessoa {
		{nome: "Ana", idade: 54, profissao: "Analista"},
		{nome: "Bruno", idade: 42, profissao: "Desenvolvedor"},
		{nome: "Carlos", idade: 51, profissao: "Analista"},
	}

	// Dois métodos de "atribuir o valor" a struct, o de cima através da função, ou através do Println.
	fmt.Println(p1)
	fmt.Println(pessoa{"João", 48, "Desenvolvedor"})
}
