package main

import "fmt"

func main() {

	// OPERADOR = && (E lógico)
	// O programa só imprime "Pode dirigir", se as duas condições forem verdadeiras
	idade := 20
	carteiraValida := true

	if (idade >= 18 && carteiraValida) {
		fmt.Println("Pode dirigir")
	} else {
		fmt.Println("Não pode dirigir")
	}

	// OPERADOR = || (OU lógico)
	// Aqui basta apenas uma das condições ser verdadeira para o if ser executado
	saldo := 0
	limiteCredito := 500

	if (saldo > 0 || limiteCredito > 0) {
		fmt.Println("Compra aprovada")
	} else {
		fmt.Println("Compra negada")
	}


	// OPERADOR = ! (NÃO lógico)
	// O operador ! inverte o valor. No caso <!false> se torna <true> e o <!true> se torna <false>
	usuarioAtivo := false

	if (!usuarioAtivo) {
		fmt.Println("Usuário inativo. Acesso negado")
	}
}